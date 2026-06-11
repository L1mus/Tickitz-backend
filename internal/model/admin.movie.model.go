package model

import "time"

type AdminMovie struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Duration    string     `json:"duration"`
	Poster      *string    `json:"poster"`
	ReleaseDate *time.Time `json:"release_date"`
	Synopsis    *string    `json:"synopsis"`
	Category    *string    `json:"category"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type AdminMovieListItem struct {
	ID          int
	Title       string
	Poster      string
	ReleaseDate *time.Time
	Duration    *string
	Genres      string
}
