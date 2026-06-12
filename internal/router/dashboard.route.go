package router

import (
	"github.com/L1mus/Tickitz-backend/internal/controller"
	"github.com/L1mus/Tickitz-backend/internal/repository"
	"github.com/L1mus/Tickitz-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func DashboardRouter(router *gin.RouterGroup, db *pgxpool.Pool, rdb *redis.Client) {
	dashboardRepo := repository.NewDashboardRepository(db)
	dashboardService := service.NewDashboardService(dashboardRepo)
	dashboardController := controller.NewDashboardController(dashboardService)

	dash := router.Group("/admin")
	dash.GET("/sales-chart", dashboardController.GetSalesChart)
	dash.GET("/ticket-sales", dashboardController.GetTicketSales)
}
