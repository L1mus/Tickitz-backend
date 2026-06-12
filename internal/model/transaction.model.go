package model

import "time"

type PaymentMethod struct {
	ID   int
	Name string
	Logo string
}

type BookingSummary struct {
	BookingID    int       `db:"booking_id"`
	UserID       int       `db:"user_id"`
	MovieTitle   string    `db:"movie_title"`
	Category     string    `db:"category"`
	CinemaName   string    `db:"cinema_name"`
	ShowDate     time.Time `db:"show_date"`
	ShowTime     string    `db:"show_time"`
	TicketPrice  int       `db:"ticket_price"`
	Quantity     int       `db:"quantity"`
	TotalPayment int       `db:"total_payment"`
	StatusPaid   string    `db:"status_paid"`
}

type BookedSeat struct {
	Label string
}

type TransactionModal struct {
	TransactionID   int
	VirtualRek      int64
	TotalPrice      int
	Status          string
	PaymentDeadline time.Time
}

type TicketResult struct {
	QRCode        string
	TotalPrice    int
	PaymentStatus string
	MovieTitle    string
	Category      string
	ShowDate      time.Time
	ShowTime      string
	TicketCount   int
	SeatLabels    string
}
