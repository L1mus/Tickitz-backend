package service

import (
	"context"
	"fmt"

	"github.com/L1mus/Tickitz-backend/internal/dto"
	"github.com/L1mus/Tickitz-backend/internal/repository"
)

type DashboardService struct {
	repo *repository.DashboardRepository
}

func NewDashboardService(repo *repository.DashboardRepository) *DashboardService {
	return &DashboardService{repo: repo}
}

func (s *DashboardService) GetSalesChart(ctx context.Context, filter dto.SalesChartFilter) ([]dto.SalesChartResponse, error) {

	data, err := s.repo.GetSalesChartData(ctx, filter)

	if err != nil {
		return nil, fmt.Errorf("dashboard_service - GetSalesChart: failed to fetch chart data: %w", err)
	}

	return data, nil
}

func (s *DashboardService) GetTicketSales(ctx context.Context, filter dto.TicketSalesFilter) ([]dto.TicketSalesResponse, error) {

	data, err := s.repo.GetTicketSalesData(ctx, filter)

	if err != nil {
		return nil, fmt.Errorf("dashboard_service - GetTicketSales: failed to fetch ticket sales data: %w", err)
	}

	return data, nil
}
