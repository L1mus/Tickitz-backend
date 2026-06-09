package controller

import (
	"errors"
	"log"
	"net/http"

	apperror "github.com/L1mus/Tickitz-backend/internal/appError"
	"github.com/L1mus/Tickitz-backend/internal/dto"
	"github.com/L1mus/Tickitz-backend/internal/response"
	"github.com/L1mus/Tickitz-backend/internal/service"
	"github.com/L1mus/Tickitz-backend/pkg"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type OrderController struct {
	orderService *service.OrderService
}

func NewOrderController(orderService *service.OrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

func (c *OrderController) GetSeats(ctx *gin.Context) {
	//_, exist := ctx.Get("claims")
	//if !exist {
	//	response.Error(ctx, http.StatusUnauthorized, "unauthorized")
	//	return
	//}
	var payload dto.OrderSeatRequest
	if err := ctx.ShouldBindWith(&payload, binding.Query); err != nil {
		log.Println(payload.ShowtimeId)
		response.Error(ctx, http.StatusBadRequest, "bad request")
		return
	}
	res, err := c.orderService.GetSeats(ctx.Request.Context(), payload.ShowtimeId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "internal server error")
	}
	response.Success(ctx, http.StatusOK, "get all information Success", res)
}

func (c *OrderController) CreateBooking(ctx *gin.Context) {
	token, _ := ctx.Get("claims")
	log.Println(token)
	claims := token.(pkg.Claims)

	var payload dto.CreateBookingRequest
	if err := ctx.ShouldBindWith(&payload, binding.JSON); err != nil {
		response.Error(ctx, http.StatusBadRequest, "bad request")
		return
	}
	res, err := c.orderService.CreateBooking(ctx.Request.Context(), payload, claims.Id)
	if err != nil {
		if errors.Is(err, apperror.InvalidSeatsInput) {
			response.Error(ctx, http.StatusNotAcceptable, err.Error())
			return
		}
		if errors.Is(err, apperror.InvalidQuantity) {
			response.Error(ctx, http.StatusNotAcceptable, err.Error())
			return
		}
		if errors.Is(err, apperror.SeatsUnavailable) {
			response.Error(ctx, http.StatusNotAcceptable, err.Error())
			return
		}
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(ctx, http.StatusOK, "success create booking", res)
}
