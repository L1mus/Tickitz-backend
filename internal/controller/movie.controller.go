package controller

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/L1mus/Tickitz-backend/internal/dto"
	"github.com/L1mus/Tickitz-backend/internal/response"
	"github.com/L1mus/Tickitz-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type MovieController struct {
	movieService *service.MovieService
}

func NewMovieController(movieService *service.MovieService) *MovieController {
	return &MovieController{
		movieService: movieService,
	}
}

func (c *MovieController) GetMovieDetail(ctx *gin.Context) {
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)
	log.Println(id)
	res, err := c.movieService.GetMovieDetail(ctx, id)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "Get detail movie success", res)
}

func (c *MovieController) GetShowtimeFilter(ctx *gin.Context) {
	var payload dto.ShowtimeFilterRequest
	idString := ctx.Param("id")
	id, _ := strconv.Atoi(idString)
	log.Println(payload)
	if err := ctx.ShouldBindWith(&payload, binding.JSON); err != nil {
		if strings.Contains(err.Error(), "required") {
			response.Error(ctx, http.StatusBadRequest, "must be filled")
			return
		}
		response.Error(ctx, http.StatusBadRequest, "bad request")
		return
	}
	res, err := c.movieService.GetShowtimeFilter(ctx, id, payload.Date, payload.City, payload.ShowTime)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "internal server error")
		return
	}
	response.Success(ctx, http.StatusOK, "Get data success", res)
}
