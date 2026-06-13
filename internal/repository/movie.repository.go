package repository

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/L1mus/Tickitz-backend/internal/dto"
	"github.com/L1mus/Tickitz-backend/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MovieRepository struct {
	db *pgxpool.Pool
}

func NewMovieRepository(db *pgxpool.Pool) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r *MovieRepository) GetMovieDetail(ctx context.Context, movieID int) (*model.MovieDetail, error) {
	const q = `
		SELECT m.id, m.title, COALESCE(m.poster, '') AS poster, m.release_date, CAST(m.duration AS TEXT) AS duration, COALESCE(m.synopsis, '') AS synopsis, COALESCE(m.category, '') AS category, STRING_AGG(DISTINCT d.name, ', ') AS directors
		FROM movies m
		LEFT JOIN movie_directors md ON m.id = md.movie_id
		LEFT JOIN directors d ON md.director_id = d.id
		WHERE m.id = $1
		GROUP BY
			m.id, m.title, m.poster, m.release_date,
			m.duration, m.synopsis, m.category`

	row := r.db.QueryRow(ctx, q, movieID)

	var movie model.MovieDetail
	err := row.Scan(&movie.ID, &movie.Title, &movie.Poster, &movie.ReleaseDate, &movie.Duration, &movie.Synopsis, &movie.Category, &movie.Directors)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *MovieRepository) GetMovieGenres(ctx context.Context, movieID int) ([]model.Genre, error) {
	const q = `
		SELECT g.id, g.genre
		FROM genres g
		JOIN movie_genres mg ON g.id = mg.genre_id
		WHERE mg.movie_id = $1`

	rows, err := r.db.Query(ctx, q, movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var genres []model.Genre
	for rows.Next() {
		var g model.Genre
		if err := rows.Scan(&g.ID, &g.Genre); err != nil {
			return nil, err
		}
		genres = append(genres, g)
	}
	return genres, rows.Err()
}

func (r *MovieRepository) GetMovieCasts(ctx context.Context, movieID int) ([]model.Cast, error) {
	const q = `
		SELECT c.id, c.name
		FROM casts c
		JOIN movie_casts mc ON c.id = mc.cast_id
		WHERE mc.movie_id = $1`

	rows, err := r.db.Query(ctx, q, movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var casts []model.Cast
	for rows.Next() {
		var c model.Cast
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		casts = append(casts, c)
	}
	return casts, rows.Err()
}

func (r *MovieRepository) GetShowtimesByFilter(
	ctx context.Context,
	movieID int,
	date time.Time,
	city string,
	showTime *string,
) ([]model.ShowtimeItem, error) {
	const q = `
		SELECT st.id AS showtime_id, c.id AS cinema_id, c.name AS cinema_name, COALESCE(c.logo, '') AS cinema_logo, st.date AS show_date, CAST(st.time AS TEXT) AS show_time, st.price AS ticket_price
		FROM showtimes st
		JOIN cinemas c ON st.cinema_id = c.id
		JOIN locations l ON c.location_id = l.id
		WHERE
			st.movie_id = $1
			AND st.date = $2
			AND l.city = $3
			AND (st.time = $4::TIME OR $4 IS NULL)
		ORDER BY st.time ASC`

	rows, err := r.db.Query(ctx, q, movieID, date, city, showTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []model.ShowtimeItem
	for rows.Next() {
		var s model.ShowtimeItem
		if err := rows.Scan(&s.ShowtimeID, &s.CinemaID, &s.CinemaName, &s.CinemaLogo, &s.ShowDate, &s.ShowTime, &s.TicketPrice); err != nil {
			return nil, err
		}
		items = append(items, s)
	}
	return items, rows.Err()
}

func (r *MovieRepository) GetAvailableLocations(ctx context.Context, movieID int) ([]model.Location, error) {
	const q = `
		SELECT DISTINCT l.id, l.city
		FROM locations l
		JOIN cinemas c ON l.id = c.location_id
		JOIN showtimes st ON c.id = st.cinema_id
		WHERE st.movie_id = $1
		ORDER BY l.city ASC`

	rows, err := r.db.Query(ctx, q, movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []model.Location
	for rows.Next() {
		var l model.Location
		if err := rows.Scan(&l.ID, &l.City); err != nil {
			return nil, err
		}
		locations = append(locations, l)
	}
	return locations, rows.Err()
}

func (r *MovieRepository) GetAllMovies(ctx context.Context, search, genre, status, limit, page, locationID string) ([]model.Movies, error) {
	q := `SELECT m.id, m.title, COALESCE(STRING_AGG(g.id || ':' || g.genre, ','), '') AS genre, m.poster, m.release_date 
          FROM movies m
          JOIN movie_genres mg ON mg.movie_id = m.id 
          JOIN genres g ON g.id = mg.genre_id
          WHERE 1=1`
	var args []any
	argCount := 1

	if search != "" {
		q += fmt.Sprintf(" AND LOWER(m.title) LIKE $%d", argCount)
		args = append(args, "%"+strings.ToLower(search)+"%")
		argCount++
	}

	if genre != "" {
		q += fmt.Sprintf(` AND m.id IN (
			SELECT movie_id FROM movie_genres 
			JOIN genres ON genres.id = movie_genres.genre_id 
			WHERE LOWER(genres.genre) = $%d
		)`, argCount)
		args = append(args, strings.ToLower(genre))
		argCount++
	}

	if status == "now_showing" {
		q += " AND m.release_date >= (CURRENT_DATE - INTERVAL '1 month')::date AND m.release_date <= CURRENT_DATE::date"
	} else if status == "upcoming" {
		q += " AND m.release_date > CURRENT_DATE::date"
	}
	if locationID != "" {
		q += fmt.Sprintf(` AND EXISTS (
            SELECT 1 FROM showtimes st
            JOIN cinemas c ON st.cinema_id = c.id
            WHERE st.movie_id = m.id AND c.location_id = $%d
        )`, argCount)
		args = append(args, locationID)
		argCount++
	}

	q += " GROUP BY m.id, m.title, m.poster, m.release_date"

	if status == "upcoming" {
		q += " ORDER BY m.release_date ASC, m.id ASC"
	} else if status == "now_showing" {
		q += " ORDER BY m.release_date DESC, m.id DESC"
	} else {
		q += " ORDER BY m.release_date ASC, m.id ASC"
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		pageInt = 1
	}

	limitSize := 12
	if limit != "" {
		if customLimit, err := strconv.Atoi(limit); err == nil && customLimit > 0 {
			limitSize = customLimit
		}
	}
	offsetSize := (pageInt - 1) * limitSize

	q += fmt.Sprintf(" LIMIT $%d OFFSET $%d", argCount, argCount+1)
	args = append(args, limitSize, offsetSize)

	rows, err := r.db.Query(ctx, q, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []model.Movies
	for rows.Next() {
		var m model.Movies
		var genreString string

		if err := rows.Scan(&m.Id, &m.Title, &genreString, &m.Poster, &m.ReleaseDate); err != nil {
			return nil, err
		}

		if genreString != "" {
			splitGenres := strings.Split(genreString, ",")
			for _, gPair := range splitGenres {
				gInfo := strings.Split(gPair, ":")
				if len(gInfo) == 2 {
					gID, _ := strconv.Atoi(gInfo[0])

					m.Genre = append(m.Genre, model.Genre{
						ID:    gID,
						Genre: strings.TrimSpace(gInfo[1]),
					})
				}
			}
		}
		movies = append(movies, m)
	}
	return movies, rows.Err()
}

func (r *MovieRepository) GetTotalCount(ctx context.Context, search, genre, status, locationID string) (int, error) {
	query := `SELECT COUNT(DISTINCT m.id) 
			  FROM movies m
			  LEFT JOIN movie_genres mg ON mg.movie_id = m.id 
			  LEFT JOIN genres g ON g.id = mg.genre_id
			  WHERE 1=1`

	var args []any
	argCount := 1

	if search != "" {
		query += fmt.Sprintf(" AND LOWER(m.title) LIKE $%d", argCount)
		args = append(args, "%"+strings.ToLower(search)+"%")
		argCount++
	}

	if genre != "" {
		query += fmt.Sprintf(` AND m.id IN (
			SELECT movie_id FROM movie_genres 
			JOIN genres ON genres.id = movie_genres.genre_id 
			WHERE LOWER(genres.genre) = $%d
		)`, argCount)
		args = append(args, strings.ToLower(genre))
		argCount++
	}

	if status == "now_showing" {
		query += " AND m.release_date >= (CURRENT_DATE - INTERVAL '1 month')::date AND m.release_date <= CURRENT_DATE::date"
	} else if status == "upcoming" {
		query += " AND m.release_date > CURRENT_DATE::date"
	}
	if locationID != "" {
		query += fmt.Sprintf(` AND EXISTS (
            SELECT 1 FROM showtimes st
            JOIN cinemas c ON st.cinema_id = c.id
            WHERE st.movie_id = m.id AND c.location_id = $%d
        )`, argCount)
		args = append(args, locationID)
		argCount++
	}

	var total int
	err := r.db.QueryRow(ctx, query, args...).Scan(&total)
	return total, err
}

func (r *MovieRepository) GetShowtimeDetailsByMovieID(ctx context.Context, movieID int) ([]dto.ShowtimeDetail, error) {
	query := `
			SELECT 
				s.id, s.date, s.time, s.price,
				l.city, c.id, c.name, c.logo, m.poster
			FROM showtimes s
			JOIN cinemas c ON s.cinema_id = c.id
			JOIN locations l ON c.location_id = l.id
			JOIN movies m ON s.movie_id = m.id
			WHERE s.movie_id = $1
			ORDER BY s.date ASC, s.time ASC;
		`

	rows, err := r.db.Query(ctx, query, movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var showtimes []dto.ShowtimeDetail

	for rows.Next() {
		var s dto.ShowtimeDetail
		err := rows.Scan(
			&s.ShowtimeID,
			&s.ShowDate,
			&s.ShowTime,
			&s.Price,
			&s.LocationName,
			&s.CinemaID,
			&s.CinemaName,
			&s.CinemaLogo,
			&s.MoviePoster,
		)
		if err != nil {
			return nil, err
		}
		showtimes = append(showtimes, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return showtimes, nil
}

func (r *MovieRepository) GetAllLocations(ctx context.Context) ([]model.Location, error) {
	const q = `
        SELECT id, city 
        FROM locations 
        ORDER BY city ASC`

	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []model.Location
	for rows.Next() {
		var l model.Location
		if err := rows.Scan(&l.ID, &l.City); err != nil {
			return nil, err
		}
		locations = append(locations, l)
	}
	return locations, rows.Err()
}
