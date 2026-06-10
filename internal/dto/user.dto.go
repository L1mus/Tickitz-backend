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
