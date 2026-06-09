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

func OrderRouter(router *gin.RouterGroup, db *pgxpool.Pool, rdb *redis.Client) {
	orderRouter := router.Group("/order")
	orderRouter.Use(middleware.VerifyToken)
	orderRepository := repository.NewOrderRepository()
	orderService := service.NewOrderService(orderRepository, db)
	orderController := controller.NewOrderController(orderService)

	orderRouter.GET("/seats", orderController.GetSeats)
	orderRouter.POST("/booking", orderController.CreateBooking)

}
