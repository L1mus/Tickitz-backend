package controller

import (
	"net/http"
	"strconv"

	"github.com/L1mus/Tickitz-backend/internal/dto"
	"github.com/L1mus/Tickitz-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type DashboardController struct {
	service *service.DashboardService
}

func NewDashboardController(s *service.DashboardService) *DashboardController {
	return &DashboardController{service: s}
}

// GetSalesChart godoc
// @Summary      Get Revenue Sales Chart Data
// @Tags         Admin
// @Param        movie_name query string false "Filter Judul Film"
// @Param        filter_by  query string true  "Filter rentang waktu" Enums(weekly, monthly)
// @Success      200 {array} dto.SalesChartResponse
// @Router       /admin/sales-chart [get]
func (c *DashboardController) GetSalesChart(ctx *gin.Context) {
	var filter dto.SalesChartFilter
	if err := ctx.ShouldBindQuery(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := c.service.GetSalesChart(ctx.Request.Context(), filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Success", "data": data})
}

// GetTicketSales godoc
// @Summary      Get Ticket Sales Data
// @Tags         Admin
// @Param        genre_id    query int false "Filter ID Genre"
// @Param        location_id query int false "Filter ID Lokasi"
// @Success      200 {array} dto.TicketSalesResponse
// @Router       /admin/ticket-sales [get]
func (c *DashboardController) GetTicketSales(ctx *gin.Context) {
	genreID, _ := strconv.Atoi(ctx.Query("genre_id"))
	locationID, _ := strconv.Atoi(ctx.Query("location_id"))

	filter := dto.TicketSalesFilter{
		GenreID:    genreID,
		LocationID: locationID,
	}

	data, err := c.service.GetTicketSales(ctx.Request.Context(), filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Success", "data": data})
}
