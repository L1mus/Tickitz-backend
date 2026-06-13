package dto

import "time"

type ShowtimeFilterRequest struct {
	Date     time.Time `json:"date" time_format:"2006-01-02" example:"2025-06-10T15:04:05Z"`
	City     string    `json:"city" example:"Jakarta"`
	ShowTime *string   `json:"show_time,omitempty" example:"hh:mm:ss"`
}
type CreateBookingRequest struct {
	ShowtimeID int   `json:"showtime_id" binding:"required" example:"1"`
	SeatIDs    []int `json:"seat_ids"    binding:"required,min=1" example:"[1,2,3,4]"`
	Quantity   int   `json:"quantity"    binding:"required,min=1" example:"4"`
}

type SubmitPaymentRequest struct {
	BookingID       int `json:"booking_id" binding:"required" example:"56"`
	PaymentMethodID int `json:"payment_method_id" binding:"required" example:"3"`
}

type OrderSeatRequest struct {
	ShowtimeId int `form:"showtime_id" binding:"required" example:"1"`
}

type ConfirmPaymentRequest struct {
	TransactionID int `json:"transaction_id" example:"5001" binding:"required"`
	BookingID     int `json:"booking_id" example:"1024" binding:"required"`
}
