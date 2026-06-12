package repository

import (
	"context"

	"github.com/L1mus/Tickitz-backend/internal/dto"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MovieOptionRepository interface {
	GetAllOptions(ctx context.Context) (dto.MovieOptionsResponse, error)
}

type movieOptionRepository struct {
	db *pgxpool.Pool
}

func NewMovieOptionRepository(db *pgxpool.Pool) MovieOptionRepository {
	return &movieOptionRepository{db: db}
}

func (r *movieOptionRepository) GetAllOptions(ctx context.Context) (dto.MovieOptionsResponse, error) {
	var response dto.MovieOptionsResponse

	genreRows, _ := r.db.Query(ctx, "SELECT id, genre FROM genres")
	defer genreRows.Close()
	for genreRows.Next() {
		var item dto.OptionItem
		if err := genreRows.Scan(&item.ID, &item.Name); err == nil {
			response.Genres = append(response.Genres, item)
		}
	}

	directorRows, _ := r.db.Query(ctx, "SELECT id, name FROM directors")
	defer directorRows.Close()
	for directorRows.Next() {
		var item dto.OptionItem
		if err := directorRows.Scan(&item.ID, &item.Name); err == nil {
			response.Directors = append(response.Directors, item)
		}
	}

	castsRows, _ := r.db.Query(ctx, "SELECT id, name FROM casts")
	defer castsRows.Close()
	for castsRows.Next() {
		var item dto.OptionItem
		if err := castsRows.Scan(&item.ID, &item.Name); err == nil {
			response.Casts = append(response.Casts, item)
		}
	}

	locationsRows, _ := r.db.Query(ctx, "SELECT id, city FROM locations")
	defer locationsRows.Close()
	for locationsRows.Next() {
		var item dto.OptionItem
		if err := locationsRows.Scan(&item.ID, &item.Name); err == nil {
			response.Locations = append(response.Locations, item)
		}
	}

	return response, nil
}
