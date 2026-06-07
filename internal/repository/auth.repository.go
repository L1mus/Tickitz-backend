package repository

import (
	"context"
	"errors"

	apperror "github.com/L1mus/Tickitz-backend/internal/appError"
	"github.com/L1mus/Tickitz-backend/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

//	type AuthRepository interface {
//		// Register(ctx context.Context, user *model.Users) (model.Users, error)
//		CheckEmailExist(ctx context.Context, email string) (bool, error)
//		Login(ctx context.Context, email string) (model.Users, error)
//	}
type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

// func(Ar *AuthRepository) Register(ctx context.Context, user *model.Users) (model.Users, error) {
// 	query := `INSERT INTO users (email, password, first_name, last_name, phone, photo, role, location_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW()) RETURNING id`

// 	err := Ar.db.QueryRow(ctx, query, user.Email, user.Password, user.First_Name, user.Last_Name, user.Phone, user.Photo, user.Role, user.Location.ID).Scan(&user.ID)
// 	if err != nil {
// 		return model.Users{}, err
// 	}
// 	return *user, nil
// }

func (ar *AuthRepository) CheckEmailExist(ctx context.Context, email string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`

	var exists bool
	err := ar.db.QueryRow(ctx, query, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
func (ar *AuthRepository) Login(ctx context.Context, email string) (model.Users, error) {
	query := `SELECT id, email, password, first_name, last_name, phone, photo, role, location_id FROM users WHERE email = $1`
	var user model.Users
	var locID *int
	err := ar.db.QueryRow(ctx, query, email).Scan(&user.ID, &user.Email, &user.Password, &user.First_Name, &user.Last_Name, &user.Phone, &user.Photo, &user.Role, &locID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Users{}, apperror.ErrUserNotFound
		}
		return model.Users{}, err
	}
	if locID != nil {
		user.Location = &model.Locations{
			ID: *locID,
		}
	}
	return user, nil
}
