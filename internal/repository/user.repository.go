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

func (r *UserRepository) GetUserProfile(ctx context.Context, userID int) (*model.UserProfile, error) {
	const q = `
		SELECT id,
			COALESCE(first_name, '') AS first_name,
			COALESCE(last_name, '') AS last_name,
			email,
			COALESCE(phone, '') AS phone,
			COALESCE(photo, '') AS photo,
			point,
			created_at,
			updated_at
		FROM users
		WHERE id = $1`

	row := r.db.QueryRow(ctx, q, userID)
	var up model.UserProfile
	err := row.Scan(&up.Id, &up.FirstName, &up.LastName, &up.Email, &up.Phone, &up.Photo, &up.Point, &up.Created_At, &up.Updated_At)
	if err != nil {
		return nil, err
	}
	return &up, nil
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
		RETURNING id, first_name, last_name, phone, photo;
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
