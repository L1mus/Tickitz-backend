package repository

import (
	"context"

	"github.com/L1mus/Tickitz-backend/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type TransactionDBTX interface {
	Query(context.Context, string, ...any) (pgx.Rows, error)
	QueryRow(context.Context, string, ...any) pgx.Row
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
}

type TransactionRepository struct{}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{}
}

func (r *TransactionRepository) GetAllPaymentMethods(ctx context.Context, tx TransactionDBTX) ([]model.PaymentMethod, error) {
	const q = `
		SELECT id, name, COALESCE(logo, '') AS logo
		FROM payment_methods
		ORDER BY id ASC`

	rows, err := tx.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var methods []model.PaymentMethod
	for rows.Next() {
		var pm model.PaymentMethod
		if err := rows.Scan(&pm.ID, &pm.Name, &pm.Logo); err != nil {
			return nil, err
		}
		methods = append(methods, pm)
	}
	return methods, rows.Err()
}

func (r *TransactionRepository) GetBookingSummary(ctx context.Context, tx TransactionDBTX, bookingID int) (*model.BookingSummary, error) {
	const q = `
		SELECT b.id AS booking_id, b.user_id, m.title AS movie_title, COALESCE(m.category, '') AS category, c.name AS cinema_name, st.date AS show_date,
			CAST(st.time AS TEXT) AS show_time, st.price AS ticket_price, b.quantity, (b.quantity * st.price) AS total_payment,
			CAST(b.status_paid AS TEXT) AS status_paid
		FROM bookings b
		JOIN showtimes st ON b.showtime_id = st.id
		JOIN movies m ON st.movie_id = m.id
		JOIN cinemas c ON st.cinema_id = c.id
		WHERE b.id = $1`

	var bs model.BookingSummary
	err := tx.QueryRow(ctx, q, bookingID).Scan(&bs.BookingID, &bs.UserID, &bs.MovieTitle, &bs.Category, &bs.CinemaName, &bs.ShowDate, &bs.ShowTime, &bs.TicketPrice, &bs.Quantity, &bs.TotalPayment, &bs.StatusPaid)
	if err != nil {
		return nil, err
	}
	return &bs, nil
}

func (r *TransactionRepository) GetBookedSeats(ctx context.Context, tx TransactionDBTX, bookingID int) ([]model.BookedSeat, error) {
	const q = `
		SELECT CONCAT(s.row, s.seat_number) AS label
		FROM booking_seats bs
		JOIN seats s ON bs.seat_id = s.id
		WHERE bs.booking_id = $1
		ORDER BY s.row ASC, s.seat_number ASC`

	rows, err := tx.Query(ctx, q, bookingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var seats []model.BookedSeat
	for rows.Next() {
		var s model.BookedSeat
		if err := rows.Scan(&s.Label); err != nil {
			return nil, err
		}
		seats = append(seats, s)
	}
	return seats, rows.Err()
}

func (r *TransactionRepository) CreateTransaction(ctx context.Context, tx TransactionDBTX, bookingID, paymentMethodID int, virtualRek int64, totalPrice int) (int, error) {
	const q = `
		INSERT INTO transactions
			(booking_id, payment_method_id, virtual_rek, total_price, status)
		VALUES ($1, $2, $3, $4, 'pending')
		RETURNING id`

	var transactionID int
	err := tx.QueryRow(ctx, q, bookingID, paymentMethodID, virtualRek, totalPrice).Scan(&transactionID)
	if err != nil {
		return 0, err
	}
	return transactionID, nil
}

func (r *TransactionRepository) GetTransactionModal(ctx context.Context, tx TransactionDBTX, bookingID int) (*model.TransactionModal, error) {
	const q = `
		SELECT t.id AS transaction_id, t.virtual_rek, t.total_price,
			CAST(t.status AS TEXT) AS status, b.created_at + INTERVAL '24 hours' AS payment_deadline
		FROM transactions t
		JOIN bookings b ON t.booking_id = b.id
		WHERE t.booking_id = $1
		ORDER BY t.id DESC
		LIMIT 1`

	var tm model.TransactionModal
	err := tx.QueryRow(ctx, q, bookingID).Scan(&tm.TransactionID, &tm.VirtualRek, &tm.TotalPrice, &tm.Status, &tm.PaymentDeadline)
	if err != nil {
		return nil, err
	}
	return &tm, nil
}

func (r *TransactionRepository) UpdateTransactionStatus(ctx context.Context, tx TransactionDBTX, transactionID int, qrCode string) error {
	const q = `
		UPDATE transactions
		SET
			status  = 'completed',
			qr_code = $2
		WHERE id = $1`

	_, err := tx.Exec(ctx, q, transactionID, qrCode)
	return err
}

func (r *TransactionRepository) UpdateBookingStatus(ctx context.Context, tx TransactionDBTX, bookingID int) error {
	const q = `
		UPDATE bookings
		SET
			status_paid   = 'paid',
			status_ticket = 'active',
			updated_at    = NOW()
		WHERE id = $1`

	_, err := tx.Exec(ctx, q, bookingID)
	return err
}

func (r *TransactionRepository) GetTicketResult(ctx context.Context, tx TransactionDBTX, transactionID int) (*model.TicketResult, error) {
	const q = `
		SELECT COALESCE(t.qr_code, '') AS qr_code, t.total_price,
			CAST(t.status AS TEXT) AS payment_status, m.title AS movie_title, COALESCE(m.category, '') AS category, st.date AS show_date,
			CAST(st.time AS TEXT) AS show_time, b.quantity AS ticket_count, STRING_AGG(CONCAT(s.row, s.seat_number),', '  ORDER BY s.row, s.seat_number ) AS seat_labels
		FROM transactions t
		JOIN bookings b ON t.booking_id = b.id
		JOIN showtimes st ON b.showtime_id = st.id
		JOIN movies m ON st.movie_id = m.id
		LEFT JOIN booking_seats bs ON bs.booking_id = b.id
		LEFT JOIN seats s ON bs.seat_id = s.id
		WHERE t.id = $1
		GROUP BY
			t.qr_code, t.total_price, t.status,
			m.title, m.category,
			st.date, st.time,
			b.quantity`

	var tr model.TicketResult
	err := tx.QueryRow(ctx, q, transactionID).Scan(&tr.QRCode, &tr.TotalPrice, &tr.PaymentStatus, &tr.MovieTitle, &tr.Category, &tr.ShowDate, &tr.ShowTime, &tr.TicketCount, &tr.SeatLabels)
	if err != nil {
		return nil, err
	}
	return &tr, nil
}

func (r *TransactionRepository) GetUser(ctx context.Context, tx TransactionDBTX, userID int) (*model.UserProfile, error) {
	const q = `
		SELECT id,
			COALESCE(first_name, '') AS first_name,
			COALESCE(last_name, '') AS last_name,
			email,
			COALESCE(phone, '') AS phone
		FROM users
		WHERE id = $1`

	row := tx.QueryRow(ctx, q, userID)
	var up model.UserProfile
	err := row.Scan(&up.Id, &up.FirstName, &up.LastName, &up.Email, &up.Phone)
	if err != nil {
		return nil, err
	}
	return &up, nil
}
