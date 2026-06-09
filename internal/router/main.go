package router

import (
	"net/http"

	_ "github.com/L1mus/Tickitz-backend/docs"
	"github.com/L1mus/Tickitz-backend/internal/dto"
	"github.com/L1mus/Tickitz-backend/internal/middleware"
	"github.com/L1mus/Tickitz-backend/pkg"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(router *gin.Engine, db *pgxpool.Pool, rdb *redis.Client, mailer pkg.Mailer) {
	// middleware global
	router.Use(middleware.CORSMiddleware)
	//swagger docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// router.METHOD(endpoint, callback)
	router.Static("/img", "public/img")

	routeApi := router.Group("/api")
	UserRouter(routeApi, db, rdb)
	MovieRouter(router, db, rdb)
	AuthRouter(routeApi, db, rdb, mailer)
	OrderRouter(routeApi, db, rdb)

	//fallback
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, dto.ResponseError{
			Status:  "Error",
			Message: "Invalid Route",
			Error:   "route not found",
		})
	})
}
