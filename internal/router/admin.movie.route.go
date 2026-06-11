package router

import (
	"github.com/L1mus/Tickitz-backend/internal/controller"
	"github.com/L1mus/Tickitz-backend/internal/repository"
	"github.com/L1mus/Tickitz-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func AdminMovieRouter(router *gin.RouterGroup, db *pgxpool.Pool, rdb *redis.Client) {
	movieRouter := router.Group("/admin")

	movieRepository := repository.AdminNewMovieRepository(db)
	movieService := service.AdminNewMovieService(movieRepository)
	movieController := controller.AdminNewMovieController(movieService)

	movieRouter.POST("/movies", movieController.AdminCreateMovie)
	movieRouter.GET("/movies", movieController.AdminGetMovies)
	movieRouter.PUT("/movies/:id", movieController.AdminUpdateMovie)
}
