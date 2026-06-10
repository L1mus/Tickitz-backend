package router

import (
	"github.com/L1mus/Tickitz-backend/internal/controller"
	"github.com/L1mus/Tickitz-backend/internal/middleware"
	"github.com/L1mus/Tickitz-backend/internal/repository"
	"github.com/L1mus/Tickitz-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func UserRouter(router *gin.RouterGroup, db *pgxpool.Pool, rdb *redis.Client) {
	userRouter := router.Group("/users")
	userRouter.Use(middleware.VerifyToken, middleware.CheckBlacklist(rdb))
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, rdb, db)
	userController := controller.NewUserController(userService)

	userRouter.GET("/profile", userController.GetUserProfile)
	userRouter.PATCH("/profile", userController.UpdateProfile)
	userRouter.GET("/history", userController.OrderHistory)
}
