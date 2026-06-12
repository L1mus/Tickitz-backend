package controller

import (
	"net/http"

	"github.com/L1mus/Tickitz-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type MovieOptionController struct {
	service service.MovieOptionService
}

func NewMovieOptionController(service service.MovieOptionService) *MovieOptionController {
	return &MovieOptionController{service: service}
}

func (c *MovieOptionController) GetOptions(ctx *gin.Context) {
	response, err := c.service.GetMovieOptions(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch options"})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
