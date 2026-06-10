package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	apperror "github.com/L1mus/Tickitz-backend/internal/appError"
	"github.com/L1mus/Tickitz-backend/internal/dto"
	"github.com/L1mus/Tickitz-backend/internal/response"
	"github.com/L1mus/Tickitz-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// @Summary      Login User
// @Description  Authenticate user using email and password to obtain a JWT token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.LoginRequest true "Login Payload"
// @Success      200 {object} dto.ResponseSuccess{data=dto.LoginResponse} "Login Successful"
// @Failure      401 {object} dto.ResponseError "Incorrect email or password"
// @Failure      403 {object} dto.ResponseError "Account not activated (ACCOUNT_NOT_ACTIVATED)"
// @Failure      500 {object} dto.ResponseError "Internal server error"
// @Router       /auth [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindWith(&req, binding.JSON); err != nil {
		errStr := err.Error()

		if strings.Contains(errStr, "password") || strings.Contains(errStr, "min=8") {
			response.Error(ctx, http.StatusBadRequest, apperror.ErrInvalidPassword.Error())
		} else if strings.Contains(errStr, "email") {
			response.Error(ctx, http.StatusBadRequest, apperror.ErrInvalidEmail.Error())
		} else if strings.Contains(errStr, "required") {
			response.Error(ctx, http.StatusBadRequest, apperror.ErrCredentialsRequired.Error())
		} else {
			response.Error(ctx, http.StatusBadRequest, "Invalid input data")
		}
		return
	}

	res, err := c.authService.Login(ctx.Request.Context(), req)
	if err != nil {
		if errors.Is(err, apperror.ErrInvalidCredentials) ||
			errors.Is(err, apperror.ErrUserNotFound) ||
			errors.Is(err, apperror.ErrInvalidPassword) {
			response.Error(ctx, http.StatusUnauthorized, apperror.ErrInvalidCredentials.Error())
			return
		}
		if errors.Is(err, apperror.ErrAccountNotActivated) {
			response.Error(ctx, http.StatusForbidden, apperror.ErrAccountNotActivated.Error())
			return
		}

		fmt.Println("LOG ERROR 500:", err.Error())
		response.Error(ctx, http.StatusInternalServerError, apperror.ErrInternalServer.Error())
		return
	}

	response.Success(ctx, http.StatusOK, "Login Success", res)
}

// @Summary      Register User
// @Description  Register a new user and send an OTP to their email
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.RegisterRequest true "Register Payload"
// @Success      201 {object} dto.ResponseSuccess{data=dto.RegisterResponse} "Registration Successful"
// @Failure      400 {object} dto.ResponseError "Invalid input data"
// @Failure      409 {object} dto.ResponseError "Email is already registered"
// @Failure      500 {object} dto.ResponseError "Internal server error"
// @Router       /auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var req dto.RegisterRequest
	if err := ctx.ShouldBindWith(&req, binding.JSON); err != nil {
		errStr := err.Error()

		if strings.Contains(errStr, "password") || strings.Contains(errStr, "min=8") {
			response.Error(ctx, http.StatusBadRequest, apperror.ErrInvalidPassword.Error())
		} else if strings.Contains(errStr, "email") {
			response.Error(ctx, http.StatusBadRequest, apperror.ErrInvalidEmail.Error())
		} else if strings.Contains(errStr, "required") {
			response.Error(ctx, http.StatusBadRequest, apperror.ErrCredentialsRequired.Error())
		} else {
			response.Error(ctx, http.StatusBadRequest, "Invalid input data")
		}
		return
	}

	res, err := c.authService.Register(ctx.Request.Context(), req)
	if err != nil {
		if errors.Is(err, apperror.ErrEmailRegistered) {
			response.Error(ctx, http.StatusConflict, apperror.ErrEmailRegistered.Error())
			return
		}

		fmt.Println("LOG ERROR 500:", err.Error())
		response.Error(ctx, http.StatusInternalServerError, apperror.ErrInternalServer.Error())
		return
	}

	response.Success(ctx, http.StatusCreated, "Check OTP in your email", res)
}

// @Summary      Activate Account
// @Description  Verify OTP to activate the user account
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.ActivationRequest true "Activation Payload"
// @Success      200 {object} dto.ResponseSuccess "Account activated successfully"
// @Failure      400 {object} dto.ResponseError "Invalid input or Invalid/expired OTP"
// @Failure      404 {object} dto.ResponseError "User not found
// @Failure      500 {object} dto.ResponseError "Internal server error"
// @Router       /auth/register/activate [post]
func (c *AuthController) Activate(ctx *gin.Context) {
	var req dto.ActivationRequest
	if err := ctx.ShouldBindWith(&req, binding.JSON); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid input data")
		return
	}

	err := c.authService.ActivateAccount(ctx.Request.Context(), req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, apperror.ErrOTPInvalid) || errors.Is(err, apperror.ErrOTPExpired) {
			statusCode = http.StatusBadRequest
		}
		if errors.Is(err, apperror.ErrUserNotFound) {
			response.Error(ctx, http.StatusNotFound, err.Error())
			return
		}

		response.Error(ctx, statusCode, err.Error())
		return
	}

	response.Success(ctx, http.StatusOK, "Account activated successfully", nil)
}

// @Summary      Resend OTP
// @Description  Resend activation OTP to the user's email
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        body body dto.ResendOTPRequest true "Email Payload"
// @Success      200 {object} dto.ResponseSuccess "OTP has been resent successfully"
// @Failure      400 {object} dto.ResponseError "Invalid input data"
// @Failure      404 {object} dto.ResponseError "User not found"
// @Failure      422 {object} dto.ResponseError "Account is already activated"
// @Failure      500 {object} dto.ResponseError "Failed to resend OTP"
// @Router       /auth/register/resend-otp [post]
func (c *AuthController) ResendOTP(ctx *gin.Context) {
	var body dto.ResendOTPRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid input")
		return
	}

	err := c.authService.ResendOTP(ctx.Request.Context(), body)
	if err != nil {
		if errors.Is(err, apperror.ErrAccountAlreadyActive) {
			response.Error(ctx, http.StatusUnprocessableEntity, err.Error())
			return
		}
		if errors.Is(err, apperror.ErrUserNotFound) {
			response.Error(ctx, http.StatusNotFound, apperror.ErrUserNotFound.Error())
			return
		}

		response.Error(ctx, http.StatusInternalServerError, "Failed to resend OTP")
		return
	}

	response.Success(ctx, http.StatusOK, "OTP resent successfully", nil)
}

// @Summary      Forgot Password
// @Description  Request an OTP to reset user password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.ForgotPasswordRequest true "Email Payload"
// @Success      200 {object} dto.ResponseSuccess "Reset password OTP sent to email"
// @Failure      400 {object} dto.ResponseError "Invalid input data"
// @Failure      404 {object} dto.ResponseError "User not found"
// @Failure      500 {object} dto.ResponseError "Internal server error"
// @Router       /auth/check-email [post]
func (c *AuthController) ForgotPassword(ctx *gin.Context) {
	var req dto.ForgotPasswordRequest
	if err := ctx.ShouldBindWith(&req, binding.JSON); err != nil {
		errStr := err.Error()
		if strings.Contains(errStr, "email") {
			response.Error(ctx, http.StatusBadRequest, apperror.ErrInvalidEmail.Error())
		} else {
			response.Error(ctx, http.StatusBadRequest, "Invalid input data")
		}
		return
	}

	err := c.authService.ForgotPassword(ctx.Request.Context(), req)
	if err != nil {
		if errors.Is(err, apperror.ErrUserNotFound) {
			response.Error(ctx, http.StatusNotFound, apperror.ErrUserNotFound.Error())
			return
		}
		response.Error(ctx, http.StatusInternalServerError, "Failed to process forgot password request")
		return
	}

	response.Success(ctx, http.StatusOK, "Please check your email for the reset OTP", nil)
}

// @Summary      Verify Reset OTP
// @Description  Verify the OTP sent for password reset
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.VerifyResetOTPReq true "OTP Payload"
// @Success      200 {object} dto.ResponseSuccess "OTP verified successfully"
// @Failure      400 {object} dto.ResponseError "Invalid input data or Invalid/Expired OTP"
// @Failure      500 {object} dto.ResponseError "Internal server error"
// @Router       /auth/check-email/verify-otp [post]
func (c *AuthController) VerifyResetOTP(ctx *gin.Context) {
	var req dto.VerifyResetOTPReq
	if err := ctx.ShouldBindWith(&req, binding.JSON); err != nil {
		errStr := err.Error()
		if strings.Contains(errStr, "email") {
			response.Error(ctx, http.StatusBadRequest, apperror.ErrInvalidEmail.Error())
		} else {
			response.Error(ctx, http.StatusBadRequest, "Invalid input data")
		}
		return
	}

	err := c.authService.VerifyResetOTP(ctx.Request.Context(), req)
	if err != nil {
		if errors.Is(err, apperror.ErrOTPInvalid) || errors.Is(err, apperror.ErrOTPExpired) {
			response.Error(ctx, http.StatusBadRequest, err.Error())
			return
		}
		response.Error(ctx, http.StatusInternalServerError, "Failed to verify OTP")
		return
	}

	response.Success(ctx, http.StatusOK, "OTP verified successfully. You can now reset your password.", nil)
}

// @Summary      Reset Password
// @Description  Set a new password after successful OTP verification
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.ResetPasswordReq true "New Password Payload"
// @Success      200 {object} dto.ResponseSuccess "Password has been reset successfully"
// @Failure      400 {object} dto.ResponseError "Invalid input data, Password Mismatch, or Session Expired"
// @Failure      500 {object} dto.ResponseError "Internal server error"
// @Router       /auth/check-email/verify-otp/reset [post]
func (c *AuthController) ResetPassword(ctx *gin.Context) {
	var req dto.ResetPasswordReq
	if err := ctx.ShouldBindWith(&req, binding.JSON); err != nil {
		errStr := err.Error()
		if strings.Contains(errStr, "new_password") || strings.Contains(errStr, "min=8") {
			response.Error(ctx, http.StatusBadRequest, apperror.ErrInvalidPassword.Error())
		} else if strings.Contains(errStr, "email") {
			response.Error(ctx, http.StatusBadRequest, apperror.ErrInvalidEmail.Error())
		} else {
			response.Error(ctx, http.StatusBadRequest, "Invalid input data")
		}
		return
	}

	err := c.authService.ResetPassword(ctx.Request.Context(), req)
	if err != nil {
		if errors.Is(err, apperror.ErrResetTokenExpired) {
			response.Error(ctx, http.StatusBadRequest, err.Error())
			return
		}
		if errors.Is(err, apperror.ErrInvalidPassword) {
			response.Error(ctx, http.StatusBadRequest, apperror.ErrInvalidPassword.Error())
			return
		}
		fmt.Println("error ini:", err)
		response.Error(ctx, http.StatusInternalServerError, "Failed to reset password")
		return
	}

	response.Success(ctx, http.StatusOK, "Password has been reset successfully", nil)
}
