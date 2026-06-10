package controller

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	apperror "github.com/L1mus/Tickitz-backend/internal/appError"
	"github.com/L1mus/Tickitz-backend/internal/dto"
	"github.com/L1mus/Tickitz-backend/internal/response"
	"github.com/L1mus/Tickitz-backend/internal/service"
	"github.com/L1mus/Tickitz-backend/pkg"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
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

// @Summary		Update User Profile
// @Description	Update user data profile
// @Tags		Profile
// @Accept		multipart/form-data
// @Produce		json
// @Security	ApiKeyAuth
// @Param		first_name			formData	string		false	"Update First Name"
// @Param		last_name			formData	string		false	"Update Last Name"
// @Param		phone				formData	string 		false	"Update Phone Number"
// @Param		new_password		formData	string		false	"Update New Password"
// @Param		confirm_password	formData	string	false	"Confirm New Password"
// @Param 		photo				formData 	file	false	"Update Photo Profile"
// @Success     200  {object}  dto.UserUpdateProfileRes "Update Profile Succesfully"
// @Failure     400  {object}  dto.ResponseError "Invalid Input Data"
// @Failure     401  {object}  dto.ResponseError "Unauthorized: Token not exist / Format token invalid"
// @Failure     422  {object}  dto.ResponseError "File Too Large / Invalid File Format"
// @Failure     500  {object}  dto.ResponseError "Failed to Update Profile / Failed to Save Image"
// @Router		/users/profile [patch]
func (c *UserController) UpdateProfile(ctx *gin.Context) {
	token, exist := ctx.Get("claims")
	if !exist {
		response.Error(ctx, http.StatusUnauthorized, "Unauthorized: Token not exist")
		return
	}
	claims, ok := token.(pkg.Claims)
	if !ok {
		response.Error(ctx, http.StatusUnauthorized, "Unauthorizzed: Format token invalid")
		return
	}

	var body dto.UserUpdateProfileReq

	if err := ctx.ShouldBind(&body); err != nil {
		log.Println("Error Bind:", err.Error())
		response.Error(ctx, http.StatusBadRequest, "Invalid Input Data")
		return
	}

	var photoURL *string
	if body.Photo != nil {
		const maxUploadSize = 2048 * 2048
		if body.Photo.Size > maxUploadSize {
			response.Error(ctx, http.StatusUnprocessableEntity, "File Too Large, Max. Size is 2MB")
			return
		}

		ext := strings.ToLower(filepath.Ext(body.Photo.Filename))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
			response.Error(ctx, http.StatusUnprocessableEntity, "Invalid File Format. JPG, JPEG or PNG Allowed")
			return
		}

		filename := fmt.Sprintf("user_%d_%d%s", claims.Id, time.Now().UnixNano(), ext)
		dst := filepath.Join("public", "img", "profiles", filename)

		if err := ctx.SaveUploadedFile(body.Photo, dst); err != nil {
			response.Error(ctx, http.StatusInternalServerError, "Failed to Save Image")
			return
		}

		generatedURL := "/img/profiles/" + filename
		photoURL = &generatedURL
	}

	res, err := c.userService.UpdateProfile(ctx.Request.Context(), claims.Id, body, photoURL)

	if err != nil {
		if err.Error() == "Confirm Password Does Not Match New Password" {
			response.Error(ctx, http.StatusBadRequest, err.Error())
			return
		}

		log.Println("Error Bind:", err.Error())
		response.Error(ctx, http.StatusInternalServerError, "Failed to Update Profile")
		return
	}

	response.Success(ctx, http.StatusOK, "Update Profile Succesfully", res)
}

// @Summary	Get User Order History
// @Description	Get list of Order History for logged-in User
// @Tags	Users
// @Accept	json
// @Produce json
// @Security	ApiKeyAuth
// @Success 200 {object} []dto.OrderHistoryRes "Get Order History Succesfully"
// @Failure     401  {object}  dto.ResponseError "Unauthorized: Token not exist / Format token invalid"
// @Failure     500  {object}  dto.ResponseError "Internal Server Error"
// @Router	/users/history [get]
func (c *UserController) OrderHistory(ctx *gin.Context) {
	token, exist := ctx.Get("claims")
	if !exist {
		response.Error(ctx, http.StatusUnauthorized, "Unauthorized: Token not exist")
		return
	}
	claims, ok := token.(*pkg.Claims)
	if !ok {
		log.Println("cek: ", claims)
		response.Error(ctx, http.StatusUnauthorized, "Unauthorizzed: Format token invalid")
		return
	}

	history, err := c.userService.GetOrderHistory(ctx.Request.Context(), claims.Id)

	if err != nil {
		if errors.Is(err, apperror.ErrUserNotFound) {
			response.Error(ctx, http.StatusUnauthorized, apperror.ErrInvalidCredentials.Error())
			return
		}

		fmt.Println("Log error: 500", err.Error())
		response.Error(ctx, http.StatusInternalServerError, apperror.ErrInternalServer.Error())
		return
	}

	response.Success(ctx, http.StatusOK, "Get Order History Succesfully", history)
}

// @Summary		Get Detail Information
// @Description	Get detailed information for specific order history by user logged-in
// @Tags		Users
// @Accept		json
// @Produce		json
// @Security	ApiKeyAuth
// @Param		id	path	int		true	"Booking Id"
// @Success		200 {object}	dto.DetailInformationRes "Get Detail Information User Succesfully"
// @Failure     400  {object}  dto.ResponseError "Invalid Booking Id Format"
// @Failure     401  {object}  dto.ResponseError "Unauthorized: Token not exist / Format token invalid"
// @Failure     500  {object}  dto.ResponseError "Internal Server Error"
// @Router		/users/history/{id}/detail [get]
func (c *UserController) DetailInformation(ctx *gin.Context) {
	token, exist := ctx.Get("claims")
	if !exist {
		response.Error(ctx, http.StatusUnauthorized, "Unauthorized: Token not exist")
		return
	}
	claims, ok := token.(*pkg.Claims)
	if !ok {
		log.Println("cek: ", claims)
		response.Error(ctx, http.StatusUnauthorized, "Unauthorizzed: Format token invalid")
		return
	}

	bookingIdStr := ctx.Param("id")
	bookingId, err := strconv.Atoi(bookingIdStr)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Invalid Booking Id Format")
		return
	}

	detail, err := c.userService.GetInformationDetail(ctx.Request.Context(), bookingId, claims.Id)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		response.Error(ctx, http.StatusInternalServerError, "Failed to Get Detail Information user")
		return
	}

	response.Success(ctx, http.StatusOK, "Get Detail Information User Succesfully", detail)
}
