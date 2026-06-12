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

func TransactionRouter(router *gin.RouterGroup, db *pgxpool.Pool, rdb *redis.Client) {
	transactionRouter := router.Group("/transactions")
	transactionRouter.Use(middleware.VerifyToken)
	transactionRepo := repository.NewTransactionRepository()
	orderRepo := repository.NewOrderRepository()
	transactionService := service.NewTransactionService(transactionRepo, orderRepo, db)
	transactionController := controller.NewTransactionController(transactionService)

	transactionRouter.GET("/payment/:id", transactionController.GetPaymentInformation)
	transactionRouter.POST("/submit", transactionController.SubmitPayment)
	transactionRouter.POST("/confirm", transactionController.ConfirmPayment)
	transactionRouter.GET("/ticket/:id", transactionController.GetResultTicket)
	transactionRouter.GET("/qr/:id", transactionController.GetQrCodeImage)
	transactionRouter.POST("/checkout", transactionController.TryCheckOutDoku)
	transactionRouter.POST("/doku-callback", transactionController.TryCallback)
}
