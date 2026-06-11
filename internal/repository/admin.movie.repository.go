package repository

import (
	"context"

	"github.com/L1mus/Tickitz-backend/internal/dto"
	"github.com/L1mus/Tickitz-backend/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AdminMovieRepository struct {
	db *pgxpool.Pool
}

func AdminNewMovieRepository(db *pgxpool.Pool) *AdminMovieRepository {
	return &AdminMovieRepository{db: db}
}

func (r *AdminMovieRepository) AdminGetMovieList(ctx context.Context, search string, month int, year int, offset int, limit int) ([]model.AdminMovieListItem, int, error) {
	countQuery := `
		SELECT COUNT(id) 
		FROM movies 
		WHERE ($1 = '' OR LOWER(title) LIKE LOWER($1))
		  AND ($2 = 0 OR EXTRACT(MONTH FROM realase_data) = $2)
		  AND ($3 = 0 OR EXTRACT(YEAR FROM realase_data) = $3)`

	searchPattern := ""
	if search != "" {
		searchPattern = "%" + search + "%"
	}

	var totalData int
	err := r.db.QueryRow(ctx, countQuery, searchPattern, month, year).Scan(&totalData)
	if err != nil {
		return nil, 0, err
	}

	dataQuery := `
		SELECT 
			m.id, 
			m.title, 
			COALESCE(m.poster, '') AS poster, 
			m.realase_data, 
			CAST(m.duration AS TEXT) AS duration, 
			COALESCE(STRING_AGG(g.genre, ', '), '') AS genres
		FROM movies m
		LEFT JOIN movie_genres mg ON m.id = mg.movie_id
		LEFT JOIN genres g ON mg.genre_id = g.id
		WHERE ($1 = '' OR LOWER(m.title) LIKE LOWER($1))
		  AND ($2 = 0 OR EXTRACT(MONTH FROM m.realase_data) = $2)
		  AND ($3 = 0 OR EXTRACT(YEAR FROM m.realase_data) = $3)
		GROUP BY m.id, m.title, m.poster, m.realase_data, m.duration
		ORDER BY m.id DESC
		LIMIT $4 OFFSET $5`

	rows, err := r.db.Query(ctx, dataQuery, searchPattern, month, year, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var movies []model.AdminMovieListItem
	for rows.Next() {
		var m model.AdminMovieListItem
		err := rows.Scan(&m.ID, &m.Title, &m.Poster, &m.ReleaseDate, &m.Duration, &m.Genres)
		if err != nil {
			return nil, 0, err
		}
		movies = append(movies, m)
	}

	if movies == nil {
		movies = []model.AdminMovieListItem{}
	}

	return movies, totalData, nil
}

func (r *AdminMovieRepository) AdminCreateMovie(
	ctx context.Context,
	movie *model.AdminMovie,
	genreIDs []int,
	castIDs []int,
	directorIDs []int,
	locationIDs []int,
	dates []string,
	times []string,
) (int, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback(ctx)

	var movieID int
	movieQuery := `
		INSERT INTO movies (title, poster, realase_data, duration, synopsis, category)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	err = tx.QueryRow(ctx, movieQuery, movie.Title, movie.Poster, movie.ReleaseDate, movie.Duration, movie.Synopsis, "").Scan(&movieID)
	if err != nil {
		return 0, err
	}

	for _, genreID := range genreIDs {
		genreQuery := `INSERT INTO movie_genres (movie_id, genre_id) VALUES ($1, $2)`
		_, err = tx.Exec(ctx, genreQuery, movieID, genreID)
		if err != nil {
			return 0, err
		}
	}

	for _, castID := range castIDs {
		castQuery := `INSERT INTO movie_casts (movie_id, cast_id) VALUES ($1, $2)`
		_, err = tx.Exec(ctx, castQuery, movieID, castID)
		if err != nil {
			return 0, err
		}
	}

	for _, directorID := range directorIDs {
		directorQuery := `INSERT INTO movie_directors (movie_id, director_id) VALUES ($1, $2)`
		_, err = tx.Exec(ctx, directorQuery, movieID, directorID)
		if err != nil {
			return 0, err
		}
	}

	cinemaRows, err := tx.Query(ctx, `SELECT id FROM cinemas WHERE location_id = ANY($1)`, locationIDs)
	if err != nil {
		return 0, err
	}

	var cinemaIDs []int
	for cinemaRows.Next() {
		var cid int
		if err := cinemaRows.Scan(&cid); err == nil {
			cinemaIDs = append(cinemaIDs, cid)
		}
	}
	cinemaRows.Close()

	defaultPrice := 50000
	showtimeQuery := `INSERT INTO showtimes (movie_id, cinema_id, date, time, price) VALUES ($1, $2, $3, $4, $5)`

	for _, cinemaID := range cinemaIDs {
		for _, d := range dates {
			for _, t := range times {
				_, err = tx.Exec(ctx, showtimeQuery, movieID, cinemaID, d, t, defaultPrice)
				if err != nil {
					return 0, err
				}
			}
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, err
	}

	return movieID, nil
}

// Edit Movie
func (r *AdminMovieRepository) AdminGetMovieByID(ctx context.Context, id int) (*model.AdminMovie, error) {
	var movie model.AdminMovie
	query := `SELECT id, title, CAST(duration AS TEXT), poster, realase_data, synopsis FROM movies WHERE id = $1`

	err := r.db.QueryRow(ctx, query, id).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Duration,
		&movie.Poster,
		&movie.ReleaseDate,
		&movie.Synopsis,
	)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *AdminMovieRepository) AdminUpdateMovieFull(ctx context.Context, id int, movie *model.AdminMovie, req dto.AdminEditMovieRequest) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	updateQuery := `UPDATE movies SET title=$1, poster=$2, realase_data=$3, duration=$4, synopsis=$5 WHERE id=$6`
	_, err = tx.Exec(ctx, updateQuery, movie.Title, movie.Poster, movie.ReleaseDate, movie.Duration, movie.Synopsis, id)
	if err != nil {
		return err
	}

	if len(req.GenreIDs) > 0 {
		_, err = tx.Exec(ctx, `DELETE FROM movie_genres WHERE movie_id = $1`, id)
		if err != nil {
			return err
		}
		for _, gid := range req.GenreIDs {
			_, err = tx.Exec(ctx, `INSERT INTO movie_genres (movie_id, genre_id) VALUES ($1, $2)`, id, gid)
			if err != nil {
				return err
			}
		}
	}

	if len(req.CastIDs) > 0 {
		_, err = tx.Exec(ctx, `DELETE FROM movie_casts WHERE movie_id = $1`, id)
		if err != nil {
			return err
		}
		for _, cid := range req.CastIDs {
			_, err = tx.Exec(ctx, `INSERT INTO movie_casts (movie_id, cast_id) VALUES ($1, $2)`, id, cid)
			if err != nil {
				return err
			}
		}
	}

	if len(req.DirectorIDs) > 0 {
		_, err = tx.Exec(ctx, `DELETE FROM movie_directors WHERE movie_id = $1`, id)
		if err != nil {
			return err
		}
		for _, did := range req.DirectorIDs {
			_, err = tx.Exec(ctx, `INSERT INTO movie_directors (movie_id, director_id) VALUES ($1, $2)`, id, did)
			if err != nil {
				return err
			}
		}
	}

	if len(req.LocationIDs) > 0 && len(req.Dates) > 0 && len(req.Times) > 0 {

		_, err = tx.Exec(ctx, `DELETE FROM showtimes WHERE movie_id = $1`, id)
		if err != nil {
			return err
		}

		rows, err := tx.Query(ctx, `SELECT id FROM cinemas WHERE location_id = ANY($1)`, req.LocationIDs)
		if err != nil {
			return err
		}

		var cinemaIDs []int
		for rows.Next() {
			var cid int
			if err := rows.Scan(&cid); err == nil {
				cinemaIDs = append(cinemaIDs, cid)
			}
		}
		rows.Close()

		for _, cID := range cinemaIDs {
			for _, d := range req.Dates {
				for _, t := range req.Times {
					_, err = tx.Exec(ctx, `INSERT INTO showtimes (movie_id, cinema_id, date, time, price) VALUES ($1, $2, $3, $4, 50000)`, id, cID, d, t)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
