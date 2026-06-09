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
	ShowtimeID  int       `json:"showtime_id"`
	CinemaID    int       `json:"cinema_id"`
	CinemaName  string    `json:"cinema_name"`
	CinemaLogo  string    `json:"cinema_logo"`
	ShowDate    time.Time `json:"show_date"`
	ShowTime    string    `json:"show_time"`
	TicketPrice int       `json:"ticket_price"`
}
type LocationDTO struct {
	ID   int    `json:"id"`
	City string `json:"city"`
}
type MovieDetailResponse struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Poster      string     `json:"poster"`
	ReleaseDate *time.Time `json:"release_date"`
	Duration    *string    `json:"duration"`
	Synopsis    string     `json:"synopsis"`
	Category    string     `json:"category"`
	Directors   string     `json:"directors"`
	Genres      []GenreDTO `json:"genres"`
	Casts       []CastDTO  `json:"casts"`
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
