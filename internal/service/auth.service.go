package service

import (
	"context"
	"errors"
	"fmt"

	apperror "github.com/L1mus/Tickitz-backend/internal/appError"
	"github.com/L1mus/Tickitz-backend/internal/dto"
	"github.com/L1mus/Tickitz-backend/internal/repository"
	"github.com/L1mus/Tickitz-backend/pkg"
	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	authRepo *repository.AuthRepository
	rdb *redis.Client
}

func NewAuthService(authRepo *repository.AuthRepository , rdb *redis.Client) *AuthService {
	return &AuthService{
		authRepo: authRepo,
		rdb:rdb,
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

	user, err := as.authRepo.Login(ctx, req.Email)
	if err != nil {
		if errors.Is(err, apperror.ErrUserNotFound) {
			return dto.LoginResponse{}, apperror.ErrUserNotFound
		}
		return dto.LoginResponse{}, apperror.ErrInternalServer
	}

	hashCfg := &pkg.HashConfig{}
	if err := hashCfg.Compare(req.Password, user.Password); err != nil {
		return dto.LoginResponse{}, apperror.ErrInvalidPassword
	}

	fullName := user.First_Name
	if fullName == "" {
		fullName = user.Email
	} else if user.Last_Name != "" {
		fullName = fmt.Sprintf("%s %s", fullName, user.Last_Name)
	}

	claims := pkg.NewClaims(int(user.ID), fullName)
	token, err := claims.GenJWT()
	if err != nil {
		return dto.LoginResponse{}, apperror.ErrInternalServer
	}

	return dto.LoginResponse{
		Message: "Login successful",
		Token:   token,
	}, nil
}