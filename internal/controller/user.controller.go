package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	apperror "github.com/L1mus/Tickitz-backend/internal/appError"
	"github.com/L1mus/Tickitz-backend/internal/response"
	"github.com/L1mus/Tickitz-backend/internal/service"
	"github.com/L1mus/Tickitz-backend/pkg"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

// @Summary		Get User Profile
// @Description	Get the profile data of the logged in user based on the jwt token
// @Tags		Users
// @Accept		json
// @Produce		json
// @Security	ApiKeyAuth
// @Success		200 {object} dto.UserProfileResponse
// @Failure		401 {object} dto.ResponseError "Unauthorized: Token not exist"
// @Failure 	500 {object} dto.ResponseError "Internal Server Error"
// @Router		/users/profile [get]
func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) GetUserProfile(ctx *gin.Context) {
	log.Println("test guys")
	token, exist := ctx.Get("claims")
	if !exist {
		response.Error(ctx, http.StatusUnauthorized, "Unauthorized: Token not exist")
		return
	}
	claims, ok := token.(pkg.Claims)
	if !ok {
		response.Error(ctx, http.StatusUnauthorized, "Unauthorizzed: Format toket invalid")
		return
	}

	profile, err := c.userService.GetProfile(ctx.Request.Context(), claims.Id)

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

	response.Success(ctx, http.StatusOK, "Get User Profile Succesfully", profile)
}
