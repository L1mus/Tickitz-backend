package repository

import (
	"context"

	"github.com/L1mus/Tickitz-backend/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserProfile(ctx context.Context, userID int) (*model.Users, error) {
	const q = `
		SELECT u.id,
			COALESCE(u.first_name, '') AS first_name,
			COALESCE(u.last_name, '') AS last_name,
			u.email,
			COALESCE(u.phone, '') AS phone,
			COALESCE(u.photo, '') AS photo,
			COALESCE(l.city, '') AS location,
			u.point,
			u.created_at,	
			u.updated_at
		FROM users u
		LEFT JOIN locations l ON u.location_id = l.id
		WHERE u.id = $1
		`

	row := r.db.QueryRow(ctx, q, userID)
	var u model.Users
	var cityName string
	err := row.Scan(&u.ID, &u.First_Name, &u.Last_Name, &u.Email, &u.Phone, &u.Photo, &cityName, &u.Point, &u.Created_At, &u.Updated_At)
	if err != nil {
		return nil, err
	}
	if cityName != "" {
		u.Location = &model.Locations{
			Name: cityName,
		}
	}
	return &u, nil
}

func (r *UserRepository) UpdateProfileById(ctx context.Context, userID int, firstName, lastName, phone, photo, hashedPassword *string) (model.UserProfile, error) {
	q := `
		UPDATE users
		SET
			first_name = COALESCE($2, first_name),
			last_name = COALESCE($3, last_name),
			phone = COALESCE($4, phone),
			photo = COALESCE($5, photo),
			password = COALESCE($6, password),
			updated_at = NOW()
		WHERE id = $1
		RETURNING 
		id,
        COALESCE(first_name, ''), 
        COALESCE(last_name, ''), 
        COALESCE(phone, ''), 
        COALESCE(photo, '')
	`

	args := []any{userID, firstName, lastName, phone, photo, hashedPassword}

	var user model.UserProfile
	err := r.db.QueryRow(ctx, q, args...).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.Photo,
	)

	if err != nil {
		return model.UserProfile{}, err
	}
	return user, nil
}

func (r *UserRepository) GetOrderHistoryById(ctx context.Context, userID int) ([]*model.OrderHistoryDetail, error) {
	const q = `
		SELECT
			b.id, b.user_id, b.showtime_id, b.status_ticket, b.status_paid, b.created_at,
			COALESCE(m.title, '') AS movie_title,
			COALESCE(c.name, '') AS cinema_name,
			COALESCE(c.logo, '') AS cinema_logo,
			s.date AS showtime_date,
			COALESCE(s.time::text, '') AS showtime_time
		FROM bookings b
		INNER JOIN showtimes s ON b.showtime_id = s.id
		INNER JOIN movies m ON s.movie_id = m.id
		INNER JOIN cinemas c ON s.cinema_id = c.id
		WHERE b.user_id = $1
		ORDER BY b.created_at DESC;
	`

	rows, err := r.db.Query(ctx, q, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []*model.OrderHistoryDetail

	for rows.Next() {
		var detail model.OrderHistoryDetail

		err := rows.Scan(
			&detail.Booking.Id,
			&detail.Booking.UserId,
			&detail.Booking.ShowtimeId,
			&detail.Booking.StatusTicket,
			&detail.Booking.StatusPaid,
			&detail.Booking.CreatedAt,
			&detail.Movie.Title,
			&detail.Cinema.Name,
			&detail.Cinema.Logo,
			&detail.Showtime.Date,
			&detail.Showtime.Time,
		)
		if err != nil {
			return nil, err
		}

		history = append(history, &detail)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return history, nil
}

func (r *UserRepository) GetDetailById(ctx context.Context, bookingID, userID int) (*model.InformationOrderDetail, error) {
	q := `
	SELECT
		b.id, b.status_ticket, b.status_paid, b.quantity, b.created_at, t.virtual_rek, t.total_price, t.qr_code, m.title AS movie_title, m.category, s.date AS showtime_date, s.time::text AS showtime_time,
		(
			SELECT string_agg(se.row || se.seat_number::text, ',')
			FROM booking_seats bs
			INNER JOIN seats se ON bs.seat_id = se.id
			WHERE bs.booking_id = b.id
		) AS seat_list
	FROM bookings b
	INNER JOIN showtimes s ON b.showtime_id = s.id
	INNER JOIN movies m ON s.movie_id = m.id
	LEFT JOIN transactions t ON b.id = t.booking_id
	WHERE b.id = $1 AND b.user_id = $2
	`
	row := r.db.QueryRow(ctx, q, bookingID, userID)

	var raw model.InformationOrderDetail
	err := row.Scan(
		&raw.BookingId, &raw.StatusTicket, &raw.StatusPaid, &raw.Quantity, &raw.CreatedAt, &raw.VirtualRek, &raw.TotalPrice, &raw.QrCode, &raw.MovieTitle, &raw.Category, &raw.ShowtimeDate, &raw.ShowtimeTime, &raw.SeatList,
	)
	if err != nil {
		return nil, err
	}
	return &raw, nil
}
