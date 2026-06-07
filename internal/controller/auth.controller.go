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
// @Description  Melakukan autentikasi menggunakan email dan password untuk mendapatkan token JWT
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.LoginRequest true "Payload Login"
// @Success      200 {object} dto.ResponseSuccess{data=dto.LoginResponse} "Login Berhasil"
// @Failure      400 {object} dto.ResponseError "Format input tidak valid"
// @Failure      401 {object} dto.ResponseError "Email atau password salah"
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

		fmt.Println("LOG ERROR 500:", err.Error())
		response.Error(ctx, http.StatusInternalServerError, apperror.ErrInternalServer.Error())
		return
	}

	response.Success(ctx, http.StatusOK, "Login Success", res)
}
