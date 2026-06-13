package dto

import (
	"mime/multipart"
)

type ProfileRequest struct {
	FirstName *string `form:"firstname" json:"firstname"`
	LastName  *string `form:"lastname" json:"lastname"`
	Phone     *string `form:"phone" json:"phone"`
}

type ChangePasswordRequest struct {
	NewPassword     string `form:"new_password" json:"new_password" binding:"required,min=6"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required,min=6"`
}

type UserProfileResponse struct {
	Id        int     `json:"id"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	FullName  *string `json:"full_name"`
	Email     string  `json:"email"`
	Photo     *string `json:"photo"`
	Phone     *string `json:"phone"`
	Point     *int    `json:"point"`
	Location  string  `json:"location"`
}

type UserUpdateProfileReq struct {
	FirstName       *string               `form:"first_name"`
	LastName        *string               `form:"last_name"`
	Phone           *string               `form:"phone"`
	Photo           *multipart.FileHeader `form:"photo"`
	NewPassword     *string               `form:"new_password" binding:"omitempty,min=6"`
	ConfirmPassword *string               `form:"confirm_password" binding:"omitempty"`
}

type UserUpdateProfileRes struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Photo     string `json:"photo"`
}

type OrderHistoryRes struct {
	BookingId    int    `json:"booking_id"`
	MovieTitle   string `json:"movie_title"`
	CinemaName   string `json:"cinema_name"`
	CinemaLogo   string `json:"cinema_logo"`
	Showtime     string `json:"showtime"`
	StatusTicket string `json:"status_ticket"`
	StatusPaid   string `json:"status_paid"`
}

type DetailInformationRes struct {
	BookingId    int      `json:"booking_id"`
	StatusTicket string   `json:"ticket_status"`
	StatusPaid   string   `json:"payment_status"`
	TotalPrice   int      `json:"total_price"`
	VirtualRek   int      `json:"virtua_account,omitempty"`
	DueDate      string   `json:"due_date,omitempty"`
	QrCode       string   `json:"qr_code,omitempty"`
	MovieTitle   string   `json:"movie_title,omitempty"`
	Category     string   `json:"category,omitempty"`
	ShowtimeTime string   `json:"showtime_time,omitempty"`
	ShowtimeDate string   `json:"showtime_date,omitempty"`
	Quantity     int      `json:"quantity,omitempty"`
	Seats        []string `json:"seats,omitempty"`
}
