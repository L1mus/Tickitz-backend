package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	apperror "github.com/L1mus/Tickitz-backend/internal/appError"
	"github.com/L1mus/Tickitz-backend/internal/cache"
	"github.com/L1mus/Tickitz-backend/internal/dto"
	"github.com/L1mus/Tickitz-backend/internal/model"
	"github.com/L1mus/Tickitz-backend/internal/repository"
	"github.com/L1mus/Tickitz-backend/pkg"
	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	authRepo repository.AuthRepository
	rdb      *redis.Client
	mailer   pkg.Mailer
}

func NewAuthService(authRepo repository.AuthRepository, rdb *redis.Client, mailer pkg.Mailer) *AuthService {
	return &AuthService{
		authRepo: authRepo,
		rdb:      rdb,
		mailer:   mailer,
	}
}

func (as *AuthService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	if req.Email == "" || req.Password == "" {
		return dto.LoginResponse{}, apperror.ErrInvalidCredentials
	}
	if len(req.Password) < 8 {
		return dto.LoginResponse{}, apperror.ErrInvalidPassword
	}
	if !pkg.IsValidEmail(req.Email) {
		return dto.LoginResponse{}, apperror.ErrInvalidEmail
	}

	user, err := as.authRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, apperror.ErrUserNotFound) {
			return dto.LoginResponse{}, apperror.ErrUserNotFound
		}
		return dto.LoginResponse{}, apperror.ErrInternalServer
	}
	if !user.Is_Active {
		otpCode, err := pkg.GenerateOTP()
		if err == nil {
			redisKey := "otp:register:" + req.Email
			as.rdb.Set(ctx, redisKey, otpCode, 5*time.Minute)

			as.mailer.SendOTP(req.Email, otpCode)
		}
		return dto.LoginResponse{}, apperror.ErrAccountNotActivated
	}

	hashCfg := &pkg.HashConfig{}
	if err := hashCfg.Compare(req.Password, user.Password); err != nil {
		return dto.LoginResponse{}, apperror.ErrInvalidPassword
	}

	var fullName string
	if user.First_Name != nil {
		fullName = *user.First_Name
	}

	if fullName == "" {
		fullName = user.Email
	} else if user.Last_Name != nil {
		fullName = fmt.Sprintf("%s %s", fullName, *user.Last_Name)
	}

	claims := pkg.NewClaims(int(user.ID), fullName, user.Role)
	token, err := claims.GenJWT()
	if err != nil {
		return dto.LoginResponse{}, apperror.ErrInternalServer
	}

	return dto.LoginResponse{
		Message: "Login successful",
		Token:   token,
		User: dto.UserDetails{
			Email: user.Email,
			Role:  user.Role,
		},
	}, nil
}

func (as *AuthService) Register(ctx context.Context, req dto.RegisterRequest) (dto.RegisterResponse, error) {
	if req.Email == "" || req.Password == "" {
		return dto.RegisterResponse{}, apperror.ErrInvalidCredentials
	}
	if len(req.Password) < 8 {
		return dto.RegisterResponse{}, apperror.ErrInvalidPassword
	}
	if !pkg.IsValidEmail(req.Email) {
		return dto.RegisterResponse{}, apperror.ErrInvalidEmail
	}

	isExist, err := as.authRepo.CheckEmailExist(ctx, req.Email)
	if err != nil {
		return dto.RegisterResponse{}, apperror.ErrInternalServer
	}
	if isExist {
		return dto.RegisterResponse{}, apperror.ErrEmailRegistered
	}

	hashCfg := &pkg.HashConfig{}
	hashCfg.UseRecommended()
	hashedPassword := hashCfg.GenHash(req.Password)

	newUser := &model.Users{
		Email:    req.Email,
		Password: hashedPassword,
		Role:     "user",
	}

	if err := as.authRepo.CreateUser(ctx, newUser); err != nil {
		return dto.RegisterResponse{}, apperror.ErrInternalServer
	}

	otpCode, err := pkg.GenerateOTP()
	if err != nil {
		return dto.RegisterResponse{}, apperror.ErrInternalServer
	}

	// save in redis
	redisKey := "otp:register:" + req.Email
	if err := as.rdb.Set(ctx, redisKey, otpCode, 5*time.Minute).Err(); err != nil {
		return dto.RegisterResponse{}, apperror.ErrInternalServer
	}
	// fmt.Printf("DEBUG: OTP Code %s is: %s\n", req.Email, otpCode)

	// send email otp
	if err := as.mailer.SendOTP(req.Email, otpCode); err != nil {
		fmt.Printf("Warning: failed to send OTP to %s: %v\n", req.Email, err)
		fmt.Println("otp: ", otpCode)
	}

	return dto.RegisterResponse{
		Message:   "Registration successful. Please check your email for the OTP.",
		Email:     req.Email,
		Is_Active: false,
	}, nil
}

func (as *AuthService) ActivateAccount(ctx context.Context, req dto.ActivationRequest) error {
	redisKey := "otp:register:" + req.Email

	// get otp in redis
	storedOTP, err := as.rdb.Get(ctx, redisKey).Result()
	if err != nil {
		if err == redis.Nil {
			return apperror.ErrOTPExpired
		}
		return apperror.ErrInternalServer
	}
	if storedOTP != req.OTP {
		return apperror.ErrOTPInvalid
	}

	if err := as.authRepo.ActivateUser(ctx, req.Email); err != nil {
		if errors.Is(err, apperror.ErrUserNotFound) {
			return apperror.ErrUserNotFound
		}
		return apperror.ErrInternalServer
	}

	_ = as.rdb.Del(ctx, redisKey)

	return nil
}

func (as *AuthService) ForgotPassword(ctx context.Context, req dto.ForgotPasswordRequest) error {
	if !pkg.IsValidEmail(req.Email) {
		return apperror.ErrInvalidEmail
	}

	isExist, err := as.authRepo.CheckEmailExist(ctx, req.Email)
	if err != nil {
		return apperror.ErrInternalServer
	}
	if !isExist {
		return apperror.ErrUserNotFound
	}

	otpCode, err := pkg.GenerateOTP()
	if err != nil {
		return apperror.ErrInternalServer
	}

	redisKey := "otp:reset:" + req.Email
	if err := as.rdb.Set(ctx, redisKey, otpCode, 5*time.Minute).Err(); err != nil {
		return apperror.ErrInternalServer
	}

	if err := as.mailer.SendResetPasswordOTP(req.Email, otpCode); err != nil {
		fmt.Printf("Warning: failed to send Reset OTP to %s: %v\n", req.Email, err)
		fmt.Println("reset_otp: ", otpCode)
	}

	return nil
}

func (as *AuthService) VerifyResetOTP(ctx context.Context, req dto.VerifyResetOTPReq) error {
	redisKey := "otp:reset:" + req.Email

	storedOTP, err := as.rdb.Get(ctx, redisKey).Result()
	if err != nil {
		if err == redis.Nil {
			return apperror.ErrOTPExpired
		}
		return apperror.ErrInternalServer
	}

	if storedOTP != req.OTP {
		return apperror.ErrOTPInvalid
	}

	_ = as.rdb.Del(ctx, redisKey)

	grantKey := "reset_granted:" + req.Email
	if err := as.rdb.Set(ctx, grantKey, "true", 5*time.Minute).Err(); err != nil {
		return apperror.ErrInternalServer
	}

	return nil
}

// 3. Menyimpan Password Baru
func (as *AuthService) ResetPassword(ctx context.Context, req dto.ResetPasswordReq) error {
	// Validasi password
	if len(req.NewPassword) < 8 {
		return apperror.ErrInvalidPassword
	}

	grantKey := "reset_granted:" + req.Email
	granted, err := as.rdb.Get(ctx, grantKey).Result()
	if err != nil || granted != "true" {
		return errors.New("unauthorized or session expired, please verify OTP again")
	}

	hashCfg := &pkg.HashConfig{}
	hashCfg.UseRecommended()
	hashedPassword := hashCfg.GenHash(req.NewPassword)

	if err := as.authRepo.ResetPassword(ctx, req.Email, hashedPassword); err != nil {
		return apperror.ErrInternalServer
	}
	_ = as.rdb.Del(ctx, grantKey)

	return nil
}

func (as *AuthService) ResendOTP(ctx context.Context, req dto.ResendOTPRequest) error {
	user, err := as.authRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, apperror.ErrUserNotFound) {
			return apperror.ErrUserNotFound
		}
		return apperror.ErrInternalServer
	}

	if user.Is_Active {
		return errors.New("account is already activated")
	}

	otpCode, err := pkg.GenerateOTP()
	if err != nil {
		return apperror.ErrInternalServer
	}

	redisKey := "otp:register:" + req.Email
	if err := as.rdb.Set(ctx, redisKey, otpCode, 5*time.Minute).Err(); err != nil {
		return apperror.ErrInternalServer
	}

	if err := as.mailer.SendOTP(req.Email, otpCode); err != nil {
		fmt.Printf("Warning: failed to resend OTP to %s: %v\n", req.Email, err)
		fmt.Println("otp: ", otpCode)

	}

	return nil
}

func (as *AuthService) Logout(ctx context.Context, token string) error {
	err := cache.SaveToBlacklist(ctx, as.rdb, token, 24*time.Hour)
	if err != nil {
		return err
	}
	return nil
}
