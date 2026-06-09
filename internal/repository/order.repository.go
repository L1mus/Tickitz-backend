package repository

import (
	"context"

	"github.com/L1mus/Tickitz-backend/internal/dto"
	"github.com/L1mus/Tickitz-backend/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type OrderDBTX interface {
	Query(context.Context, string, ...any) (pgx.Rows, error)
	QueryRow(context.Context, string, ...any) pgx.Row
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
}

type OrderRepository struct{}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (r *OrderRepository) GetShowtimeSummary(ctx context.Context, tx OrderDBTX, showtimeID int) (*model.ShowtimeSummary, error) {
	const q = `
		SELECT m.title AS movie_title, COALESCE(m.poster, '') AS movie_poster, COALESCE(m.category, '') AS category, c.id AS cinema_id, c.name AS cinema_name, COALESCE(c.logo, '') AS cinema_logo, st.date AS show_date, CAST(st.time AS TEXT) AS show_time, st.price AS ticket_price
		FROM showtimes st
		JOIN movies m ON st.movie_id = m.id
		JOIN cinemas c ON st.cinema_id = c.id
		WHERE st.id = $1`

	var data model.ShowtimeSummary
	err := tx.QueryRow(ctx, q, showtimeID).Scan(&data.MovieTitle, &data.MoviePoster, &data.Category, &data.CinemaID, &data.CinemaName, &data.CinemaLogo, &data.ShowDate, &data.ShowTime, &data.TicketPrice)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *OrderRepository) GetSeatsByShowtime(ctx context.Context, tx OrderDBTX, showtimeID, cinemaID int) ([]model.SeatRow, error) {
	const q = `
		SELECT s.id AS seat_id, s.row, s.seat_number,
			CAST(s.seat_type AS TEXT) AS seat_type,
			CASE
				WHEN bs.id IS NOT NULL THEN 'Sold'
				ELSE 'Available'
			END AS seat_status
		FROM seats s
		LEFT JOIN booking_seats bs
			ON  s.id = bs.seat_id
			AND bs.showtime_id = $1
		WHERE s.cinema_id = $2
		ORDER BY s.row ASC, s.seat_number ASC`

	rows, err := tx.Query(ctx, q, showtimeID, cinemaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seats []model.SeatRow
	for rows.Next() {
		var s model.SeatRow
		if err := rows.Scan(&s.SeatID, &s.Row, &s.SeatNumber, &s.SeatType, &s.SeatStatus); err != nil {
			return nil, err
		}
		seats = append(seats, s)
	}
	return seats, rows.Err()
}

func (r *OrderRepository) CreateBooking(ctx context.Context, tx OrderDBTX, req dto.CreateBookingRequest, userID int) (int, error) {
	const q = `
		INSERT INTO bookings
			(user_id, showtime_id, status_ticket, status_paid, quantity)
		VALUES
			($1, $2, 'active', 'not_paid', $3)	
		RETURNING id`

	var bookingID int
	err := tx.QueryRow(ctx, q, userID, req.ShowtimeID, req.Quantity).Scan(&bookingID)
	if err != nil {
		return 0, err
	}
	return bookingID, nil
}

func (r *OrderRepository) CreateBookingSeat(ctx context.Context, tx OrderDBTX, bookingID, seatID, showtimeID int) error {
	const q = `
		INSERT INTO booking_seats
			(booking_id, seat_id, showtime_id)
		VALUES
			($1, $2, $3)`

	_, err := tx.Exec(ctx, q, bookingID, seatID, showtimeID)
	return err
}
