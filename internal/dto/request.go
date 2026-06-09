package dto

import "time"

type ShowtimeFilterRequest struct {
	Date     time.Time `json:"date"`
	City     string    `json:"city"`
	ShowTime *string   `json:"show_time,omitempty"`
}

type CreateBookingRequest struct {
	ShowtimeID int   `json:"showtime_id" binding:"required"`
	SeatIDs    []int `json:"seat_ids"    binding:"required,min=1"`
	Quantity   int   `json:"quantity"    binding:"required,min=1"`
}

type SubmitPaymentRequest struct {
	PaymentMethodID int `json:"payment_method_id" binding:"required"`
}

type OrderSeatRequest struct {
	ShowtimeId int `form:"showtime_id" binding:"required"`
}
