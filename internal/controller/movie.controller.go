package controller

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"

	apperror "github.com/L1mus/Tickitz-backend/internal/appError"
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

// @Summary      Get All movies
// @Description  fetch all movie data to be displayed on the main page
// @Tags         Movies
// @Accept       json
// @Produce      json
// @Success      200 {object} dto.MoviePaginationResponse "Success to get data"
// @Failure      500 {object} dto.ResponseError "Failed to get data"
// @Router       /movies [get]
func (c *MovieController) GetAll(ctx *gin.Context) {
	searchParam := ctx.Query("search")
	genreParam := ctx.Query("genre")
	statusParam := ctx.Query("status")
	limitParam := ctx.Query("limit")
	pageParam := ctx.DefaultQuery("page", "1")

	movies, err := c.movieService.GetAllMovies(ctx.Request.Context(), searchParam, genreParam, statusParam, limitParam, pageParam)
	if err != nil {
		fmt.Println("LOG ERROR 500:", err.Error())
		response.Error(ctx, http.StatusInternalServerError, apperror.ErrInternalServer.Error())
		return
	}

	totalData, _ := c.movieService.GetTotalCount(ctx.Request.Context(), searchParam, genreParam, statusParam)

	pageInt, _ := strconv.Atoi(pageParam)
	limitInt := 12
	if limitParam != "" {
		limitInt, _ = strconv.Atoi(limitParam)
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(limitInt)))

	var nextLink, prevLink string
	baseUrl := "http://localhost:8080/api/movies"

	queryParams := ctx.Request.URL.Query()

	if pageInt < totalPage {
		queryParams.Set("page", strconv.Itoa(pageInt+1))
		nextLink = fmt.Sprintf("%s?%s", baseUrl, queryParams.Encode())
	}

	if pageInt > 1 {
		queryParams.Set("page", strconv.Itoa(pageInt-1))
		prevLink = fmt.Sprintf("%s?%s", baseUrl, queryParams.Encode())
	}

	response := dto.MoviePaginationResponse{
		Status:  "Success",
		Message: "successfully obtained the film data",
		Meta: dto.PaginationMeta{
			CurrentPage: pageInt,
			Limit:       limitInt,
			TotalData:   totalData,
			TotalPage:   totalPage,
			HasNext:     pageInt < totalPage,
			HasPrev:     pageInt > 1,
			Next:        nextLink,
			Prev:        prevLink,
		},
		Data: movies,
	}

	ctx.JSON(http.StatusOK, response)
}

// @Summary		Get Movie Detail
// @Description	Get the Movie detail data
// @Tags		Movies
// @Produce		json
// @Security	ApiKeyAuth
// @Param        id   path      int  true  "Movie ID"
// @Success		200 {object} dto.MovieDetailResponse
// @Failure 	500 {object} dto.ResponseError "Internal Server Error"
// @Router		/movies/:id [get]
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

// @Summary		Get Showtime Filter
// @Description	Get the Showtime data filtered result
// @Tags         Movies
// @Accept       json
// @Produce      json
// @Param        id       path      int                          true  "Movie ID"
// @Param        body     body      dto.ShowtimeFilterRequest    true  "Payload filter showtime"
// @Success		200 {object} dto.ShowtimeFilterResponse
// @Failure 	400 {object} dto.ResponseError "bad request" "must be filled"
// @Failure 	500 {object} dto.ResponseError "internal Server Error"
// @Router		/movies/:id/showtime [get]
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
