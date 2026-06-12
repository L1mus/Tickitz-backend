package dto

import (
	"mime/multipart"
	"time"
)

type AdminMovieQueryParams struct {
	Page   int    `form:"page,default=1"`
	Limit  int    `form:"limit,default=10"`
	Search string `form:"search"`
	Month  int    `form:"month"`
	Year   int    `form:"year"`
}

type AdminMovieListDTO struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Poster      string     `json:"poster"`
	ReleaseDate *time.Time `json:"release_date"`
	Duration    *string    `json:"duration"`
	Genres      string     `json:"genres"`
}

type AdminMovieListResponse struct {
	Movies    []AdminMovieListDTO `json:"movies"`
	TotalData int                 `json:"total_data"`
	TotalPage int                 `json:"total_page"`
	Page      int                 `json:"page"`
	Limit     int                 `json:"limit"`
}

type AdminAddMovieRequest struct {
	Title          string                `form:"title" binding:"required"`
	Poster         *multipart.FileHeader `form:"poster" binding:"required"`
	ReleaseDate    string                `form:"release_date" binding:"required"`
	DurationHour   int                   `form:"duration_hour" binding:"min=0"`
	DurationMinute int                   `form:"duration_minute" binding:"min=0,max=59"`
	Synopsis       string                `form:"synopsis" binding:"required"`
	GenreIDs       []int                 `form:"genre_ids" binding:"required"`
	CastIDs        []int                 `form:"cast_ids" binding:"required"`
	DirectorIDs    []int                 `form:"director_ids" binding:"required"`
	LocationIDs    []int                 `form:"location_ids" binding:"required"`
	Dates          []string              `form:"dates" binding:"required"`
	Times          []string              `form:"times" binding:"required"`
}

type AdminEditMovieRequest struct {
	Title          string                `form:"title"`
	Poster         *multipart.FileHeader `form:"poster"`
	ReleaseDate    string                `form:"release_date"`
	DurationHour   int                   `form:"duration_hour"`
	DurationMinute int                   `form:"duration_minute"`
	Synopsis       string                `form:"synopsis"`
	GenreIDs       []int                 `form:"genre_ids"`
	CastIDs        []int                 `form:"cast_ids"`
	DirectorIDs    []int                 `form:"director_ids"`
	LocationIDs    []int                 `form:"location_ids"`
	Dates          []string              `form:"dates"`
	Times          []string              `form:"times"`
}

type OptionItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MovieOptionsResponse struct {
	Genres    []OptionItem `json:"genres"`
	Directors []OptionItem `json:"directors"`
	Casts     []OptionItem `json:"casts"`
	Locations []OptionItem `json:"locations"`
}

// AdminResponseSuccess adalah struktur dasar untuk merespon request yang berhasil di Swagger
type AdminResponseSuccess struct {
	Status  string      `json:"status" example:"success"`
	Message string      `json:"message" example:"Operation successful"`
	Data    interface{} `json:"data"`
}

// AdminResponseError adalah struktur dasar untuk merespon request yang gagal di Swagger
type AdminResponseError struct {
	Status  string `json:"status" example:"error"`
	Message string `json:"message" example:"Error message detail"`
}
