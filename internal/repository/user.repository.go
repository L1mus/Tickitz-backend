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
