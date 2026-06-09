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
type ShowtimeFilterRequest struct {
	Date     time.Time `json:"date"`
	City     string    `json:"city"`
	ShowTime *string   `json:"show_time,omitempty"`
}
type ShowtimeFilterResponse struct {
	Showtimes []ShowtimeItemDTO `json:"showtimes"`
	Locations []LocationDTO     `json:"locations"`
}
type SeatDTO struct {
	SeatID     int    `json:"seat_id"`
	Row        string `json:"row"`
	SeatNumber int    `json:"seat_number"`
	SeatType   string `json:"seat_type"`
	SeatStatus string `json:"seat_status" binding:"required,oneof=available sold"`
}
type SeatPageResponse struct {
	MovieTitle  string    `json:"movie_title"`
	MoviePoster string    `json:"movie_poster"`
	Category    string    `json:"category"`
	CinemaName  string    `json:"cinema_name"`
	CinemaLogo  string    `json:"cinema_logo"`
	ShowDate    time.Time `json:"show_date"`
	ShowTime    string    `json:"show_time"`
	TicketPrice int       `json:"ticket_price"`
	Seats       []SeatDTO `json:"seats"`
}
type CreateBookingRequest struct {
	ShowtimeID int   `json:"showtime_id" binding:"required"`
	SeatIDs    []int `json:"seat_ids"    binding:"required,min=1"`
	Quantity   int   `json:"quantity"    binding:"required,min=1"`
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
type SubmitPaymentRequest struct {
	PaymentMethodID int `json:"payment_method_id" binding:"required"`
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
type UserProfileResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type MovieResponse struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Poster      string     `json:"poster"`
	Genres      []GenreDTO `json:"genres"`
	ReleaseDate *time.Time `json:"release_date"`
}
