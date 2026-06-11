package repository

import (
	"context"
	"fmt"

	"github.com/L1mus/Tickitz-backend/internal/dto"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DashboardRepository struct {
	db *pgxpool.Pool
}

func NewDashboardRepository(db *pgxpool.Pool) *DashboardRepository {
	return &DashboardRepository{db: db}
}

func (r *DashboardRepository) GetSalesChartData(ctx context.Context, filter dto.SalesChartFilter) ([]dto.SalesChartResponse, error) {
	results := []dto.SalesChartResponse{}
	var timeFormat string
	var groupBy string
	var dateTruncParam string

	if filter.FilterBy == "monthly" {
		timeFormat = "Mon YYYY"
		groupBy = "DATE_TRUNC('month', b.created_at)"
		dateTruncParam = "month"
	} else {
		timeFormat = `"Week "WW YYYY`
		groupBy = "DATE_TRUNC('week', b.created_at)"
		dateTruncParam = "week"
	}

	query := fmt.Sprintf(`
		SELECT
			TO_CHAR(%s, '%s') as label,
			SUM(t.total_price) as total_revenue
		FROM transactions t
		JOIN bookings b ON t.booking_id = b.id
		JOIN showtimes s ON b.showtime_id = s.id
		JOIN movies m ON s.movie_id = m.id
		WHERE t.status = 'completed'
	`, groupBy, timeFormat)

	args := []interface{}{}
	if filter.MovieName != "" {
		query += " AND m.title ILIKE $1"
		args = append(args, "%"+filter.MovieName+"%")
	}

	query += fmt.Sprintf(" GROUP BY %s, DATE_TRUNC('%s', b.created_at) ORDER BY DATE_TRUNC('%s', b.created_at) ASC", groupBy, dateTruncParam, dateTruncParam)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var res dto.SalesChartResponse
		if err := rows.Scan(&res.Label, &res.TotalRevenue); err != nil {
			return nil, err
		}
		results = append(results, res)
	}

	return results, nil
}

func (r *DashboardRepository) GetTicketSalesData(ctx context.Context, filter dto.TicketSalesFilter) ([]dto.TicketSalesResponse, error) {
	var results []dto.TicketSalesResponse

	query := `
		SELECT
			m.title as movie_title,
			SUM(b.quantity) as total_sold
		FROM bookings b
		JOIN transactions t ON t.booking_id = b.id
		JOIN showtimes s ON b.showtime_id = s.id
		JOIN movies m ON s.movie_id = m.id
		JOIN cinemas c ON s.cinema_id = c.id
		LEFT JOIN movie_genres mg ON m.id = mg.movie_id
		WHERE t.status = 'completed'
	`

	args := []interface{}{}
	argCount := 1

	if filter.GenreID != 0 {
		query += fmt.Sprintf(" AND mg.genre_id = $%d", argCount)
		args = append(args, filter.GenreID)
		argCount++
	}

	if filter.LocationID != 0 {
		query += fmt.Sprintf(" AND c.location_id = $%d", argCount)
		args = append(args, filter.LocationID)
	}

	query += " GROUP BY m.id, m.title ORDER BY total_sold DESC"

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var res dto.TicketSalesResponse
		if err := rows.Scan(&res.MovieTitle, &res.TotalSold); err != nil {
			return nil, err
		}
		results = append(results, res)
	}

	return results, nil
}
