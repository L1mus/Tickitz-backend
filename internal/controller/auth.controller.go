package controller

import (
	"errors"
	"net/http"

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

	// bad req: 400
	if err := ctx.ShouldBindWith(&req, binding.JSON); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := c.authService.Login(ctx.Request.Context(), req)
	if err != nil {
		if errors.Is(err, apperror.ErrInvalidEmail) {
			response.Error(ctx, http.StatusBadRequest, "format email tidak valid")
			return
		}
		// kredensial: 401
		if errors.Is(err, apperror.ErrInvalidCredentials) || errors.Is(err, apperror.ErrUserNotFound) || errors.Is(err, apperror.ErrInvalidPassword) {
			response.Error(ctx, http.StatusUnauthorized, "email atau password salah")
			return
		}
		
		// internal server error: 500
		response.Error(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	response.Success(ctx, http.StatusOK, res.Message, res)
}

