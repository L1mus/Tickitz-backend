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
	Genres      []GenreDTO `json:"genres" example:"[Sci-Fi, Adventure]"`
	Casts       []CastDTO  `json:"casts" example:"[Cillian Murphy, Zendaya]"`
}

type ShowtimeFilterResponse struct {
	Showtime  []ShowtimeItemDTO `json:"showtime"`
	Locations []LocationDTO     `json:"locations"`
}
type SeatDTO struct {
	SeatID     int    `json:"seat_id"`
	Row        string `json:"row"`
	SeatNumber int    `json:"seat_number"`
	SeatType   string `json:"seat_type"`
	SeatStatus string `json:"seat_status" binding:"required,oneof=available sold"`
}

type SummaryMovieDTO struct {
	MovieTitle  string    `json:"movie_title"`
	MoviePoster string    `json:"movie_poster"`
	Category    string    `json:"category"`
	CinemaName  string    `json:"cinema_name"`
	CinemaLogo  string    `json:"cinema_logo"`
	ShowDate    time.Time `json:"show_date"`
	ShowTime    string    `json:"show_time"`
	TicketPrice int       `json:"ticket_price"`
}
type SeatPageResponse struct {
	Summary SummaryMovieDTO `json:"summary"`
	Seats   []SeatDTO       `json:"seats"`
}

type CreateBookingResponse struct {
	BookingID int `json:"booking_id"`
}
type PaymentMethodDTO struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}
type PaymentPageResponse struct {
	BookingID      int                `json:"booking_id"`
	MovieTitle     string             `json:"movie_title"`
	CinemaName     string             `json:"cinema_name"`
	ShowDate       time.Time          `json:"show_date"`
	ShowTime       string             `json:"show_time"`
	Quantity       int                `json:"quantity"`
	SeatLabels     string             `json:"seat_labels"`
	TotalPrice     int                `json:"total_price"`
	FullName       string             `json:"full_name"`
	Email          string             `json:"email"`
	Phone          string             `json:"phone"`
	PaymentMethods []PaymentMethodDTO `json:"payment_methods"`
}

type TransactionModalResponse struct {
	TransactionID   int       `json:"transaction_id"`
	VirtualRek      int64     `json:"virtual_rek"`
	TotalPrice      int       `json:"total_price"`
	Status          string    `json:"status"`
	PaymentDeadline time.Time `json:"payment_deadline"`
}
type TicketResultResponse struct {
	QRCode        string    `json:"qr_code"`
	TotalPrice    int       `json:"total_price"`
	PaymentStatus string    `json:"payment_status"`
	MovieTitle    string    `json:"movie_title"`
	Category      string    `json:"category"`
	ShowDate      time.Time `json:"show_date"`
	ShowTime      string    `json:"show_time"`
	TicketCount   int       `json:"ticket_count"`
	SeatLabels    string    `json:"seat_labels"`
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
