package router

import (
	"github.com/L1mus/Tickitz-backend/internal/controller"
	"github.com/L1mus/Tickitz-backend/internal/repository"
	"github.com/L1mus/Tickitz-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func MovieRouter(router *gin.RouterGroup, db *pgxpool.Pool, rdb *redis.Client) {
	movieRouter := router.Group("/movies")

	repositryMovie := repository.NewMovieRepository(db)
	serviceMovie := service.NewMovieService(repositryMovie)
	controllerMovie := controller.NewMovieController(serviceMovie)

	movieRouter.GET("", controllerMovie.GetAll)
	movieRouter.GET("/locations", controllerMovie.GetAllLocations)
	movieRouter.GET("/:id", controllerMovie.GetMovieDetail)
	// movieRouter.GET("/:id/showtime", controllerMovie.GetShowtimeFilter)
	movieRouter.GET("/:id/showtimes", controllerMovie.GetShowtimes)
}
