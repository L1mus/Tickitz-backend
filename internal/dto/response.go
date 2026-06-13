package dto

import "time"

type GenreDTO struct {
	ID    int    `json:"id"`
	Genre string `json:"genre"`
}
type CastDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ResponseSuccess struct {
	Status  string `json:"status" example:"success"`
	Message string `json:"message" example:"Welcome, John doe"`
}

type ResponseError struct {
	Status  string `json:"status" example:"error"`
	Message string `json:"message" example:"Failed get data/internal server error"`
	Error   string `json:"error" example:"internal server error/bad request"`
}
type ShowtimeItemDTO struct {
	ShowtimeID  int       `json:"showtime_id" example:"2"`
	CinemaID    int       `json:"cinema_id" example:"1"`
	CinemaName  string    `json:"cinema_name" example:"ebv.id Grand Indonesia"`
	CinemaLogo  string    `json:"cinema_logo" example:" https://storage.tickitz.id/cinema/ebvid.png"`
	ShowDate    time.Time `json:"show_date" example:"2025-06-10"`
	ShowTime    string    `json:"show_time" example:"13:30:00"`
	TicketPrice int       `json:"ticket_price" example:"75000"`
}
type LocationDTO struct {
	ID   int    `json:"id" example:"1"`
	City string `json:"city" example:"Jakarta"`
}
type MovieDetailResponse struct {
	ID          int        `json:"id" example:"1"`
	Title       string     `json:"title" example:"Interstellar Reborn"`
	Poster      string     `json:"poster" example:" https://storage.tickitz.id/posters/interstellar_reborn.jpg"`
	ReleaseDate *time.Time `json:"release_date" example:"2025-03-15"`
	Duration    *string    `json:"duration" example:"2:49:00"`
	Synopsis    string     `json:"synopsis" example:"Seorang astronot nekat melintasi lubang cacing demi menyelamatkan umat manusia dari kehancuran bumi yang semakin tak layak huni. Perjalanan melintasi ruang dan waktu membawa konsekuensi yang tak pernah ia bayangkan sebelumnya."`
	Category    string     `json:"category" example:"13+"`
	Directors   string     `json:"directors" example:"Christopher Nolan"`
	Genres      []GenreDTO `json:"genres" `
	Casts       []CastDTO  `json:"casts" `
}

type ShowtimeFilterResponse struct {
	Showtime  []ShowtimeItemDTO `json:"showtime"`
	Locations []LocationDTO     `json:"locations"`
}
type SeatDTO struct {
	SeatID     int    `json:"seat_id" example:"145"`
	Row        string `json:"row" example:"C"`
	SeatNumber int    `json:"seat_number" example:"7"`
	SeatType   string `json:"seat_type" example:"Regular"`
	SeatStatus string `json:"seat_status" example:"available" binding:"required,oneof=available sold"`
}

type SummaryMovieDTO struct {
	MovieTitle  string    `json:"movie_title" example:"Interstellar Reborn"`
	MoviePoster string    `json:"movie_poster" example:"https://storage.tickitz.id/posters/interstellar_reborn.jpg"`
	Category    string    `json:"category" example:"13+"`
	CinemaName  string    `json:"cinema_name" example:"CGV Grand Indonesia"`
	CinemaLogo  string    `json:"cinema_logo" example:"https://storage.tickitz.id/cinema/cgv.png"`
	ShowDate    time.Time `json:"show_date" example:"2026-06-15T00:00:00Z"`
	ShowTime    string    `json:"show_time" example:"18:30"`
	TicketPrice int       `json:"ticket_price" example:"50000"`
}
type SeatPageResponse struct {
	Summary SummaryMovieDTO `json:"summary"`
	Seats   []SeatDTO       `json:"seats"`
}

type CreateBookingResponse struct {
	BookingID int `json:"booking_id" example:"1024"`
}
type PaymentMethodDTO struct {
	ID   int    `json:"id" example:"2"`
	Name string `json:"name" example:"DANA"`
	Logo string `json:"logo" example:"/img/pm/DANA.png"`
}
type PaymentPageResponse struct {
	BookingID      int                `json:"booking_id" example:"1024"`
	MovieTitle     string             `json:"movie_title" example:"Interstellar Reborn"`
	CinemaName     string             `json:"cinema_name" example:"CGV Grand Indonesia"`
	ShowDate       time.Time          `json:"show_date" example:"2026-06-15T00:00:00Z"`
	ShowTime       string             `json:"show_time" example:"18:30"`
	Quantity       int                `json:"quantity" example:"2"`
	SeatLabels     string             `json:"seat_labels" example:"C7, C8"`
	TotalPrice     int                `json:"total_price" example:"100000"`
	FullName       string             `json:"full_name" example:"Ali Mustadji"`
	Email          string             `json:"email" example:"ali.mustadji@example.com"`
	Phone          string             `json:"phone" example:"081234567890"`
	PaymentMethods []PaymentMethodDTO `json:"payment_methods"`
}
type TransactionModalResponse struct {
	TransactionID   int       `json:"transaction_id" example:"5001"`
	VirtualRek      int64     `json:"virtual_rek" example:"8856001234567890"`
	TotalPrice      int       `json:"total_price" example:"100000"`
	Status          string    `json:"status" example:"PENDING"`
	PaymentDeadline time.Time `json:"payment_deadline" example:"2026-06-13T19:47:00Z"`
}
type TicketResultResponse struct {
	QRCode        string    `json:"qr_code" example:"TICKITZ-TX-5001-SECURE-STRING"`
	TotalPrice    int       `json:"total_price" example:"100000"`
	PaymentStatus string    `json:"payment_status" example:"PAID"`
	MovieTitle    string    `json:"movie_title" example:"Interstellar Reborn"`
	Category      string    `json:"category" example:"13+"`
	ShowDate      time.Time `json:"show_date" example:"2026-06-15T00:00:00Z"`
	ShowTime      string    `json:"show_time" example:"18:30"`
	TicketCount   int       `json:"ticket_count" example:"2"`
	SeatLabels    string    `json:"seat_labels" example:"C7, C8"`
}

type MovieResponse struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Poster      string     `json:"poster"`
	Genres      []GenreDTO `json:"genres"`
	ReleaseDate *time.Time `json:"release_date"`
}

type PaginationMeta struct {
	CurrentPage int    `json:"current_page"`
	Limit       int    `json:"limit"`
	TotalData   int    `json:"total_data"`
	TotalPage   int    `json:"total_page"`
	HasNext     bool   `json:"has_next"`
	HasPrev     bool   `json:"has_prev"`
	Next        string `json:"next"`
	Prev        string `json:"prev"`
}

type MoviePaginationResponse struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Meta    PaginationMeta  `json:"meta"`
	Data    []MovieResponse `json:"data"`
}

type ShowtimeDetail struct {
	ShowtimeID   int       `json:"showtime_id"`
	ShowDate     time.Time `json:"show_date"`
	ShowTime     string    `json:"show_time"`
	Price        int       `json:"price"`
	LocationName string    `json:"location_name"`
	CinemaID     int       `json:"cinema_id"`
	CinemaName   string    `json:"cinema_name"`
	CinemaLogo   string    `json:"cinema_logo"`
	MoviePoster  string    `json:"movie_poster"`
}

type ShowtimeDetailResponse struct {
	ShowtimeID  int    `json:"showtime_id"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	Price       int    `json:"price"`
	City        string `json:"city"`
	CinemaID    int    `json:"cinema_id"`
	CinemaName  string `json:"cinema_name"`
	CinemaLogo  string `json:"cinema_logo"`
	MoviePoster string `json:"movie_poster"`
}
