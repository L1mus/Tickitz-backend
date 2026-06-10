package model

import (
	"time"
)

type Users struct {
	ID         int        `json:"id" db:"id"`
	Email      string     `json:"email" db:"email"`
	Password   string     `json:"-" db:"password"`
	First_Name *string    `json:"first_name" db:"first_name"`
	Last_Name  *string    `json:"last_name" db:"last_name"`
	Phone      *string    `json:"phone" db:"phone"`
	Photo      *string    `json:"photo" db:"photo"`
	Role       string     `json:"role" db:"role"`
	Location   *Locations `json:"location" db:"location_id"`
	Is_Active  bool       `json:"is_active" db:"isactive"`
	Created_At time.Time  `json:"created_at" db:"created_at"`
	Updated_At time.Time  `json:"updated_at" db:"updated_at"`
}

type Locations struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"city"`
}

type UserProfile struct {
	Id         int        `json:"id" db:"id"`
	FirstName  string     `json:"first_name" db:"first_name"`
	LastName   string     `json:"last_name" db:"last_name"`
	Email      string     `json:"email" db:"email"`
	Phone      string     `json:"phone" db:"phone"`
	Photo      string     `json:"photo" db:"photo"`
	Point      int        `json:"point" db:"point"`
	Created_At time.Time  `json:"created_at" db:"created_at"`
	Updated_At *time.Time `json:"updated_at" db:"updated_at"`
}

type Movie struct {
	Id    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
}

type Cinema struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Logo string `db:"logo" json:"logo"`
}

type Showtime struct {
	Id       int       `db:"id" json:"id"`
	MovieId  int       `db:"movie_id" json:"movie_id"`
	CinemaId int       `db:"cinema_id" json:"cinema_id"`
	Date     time.Time `db:"date" json:"date"`
	Time     string    `db:"time" json:"time"`
}

type Booking struct {
	Id           int       `db:"id" json:"id"`
	UserId       int       `db:"user_id" json:"user_id"`
	ShowtimeId   int       `db:"showtime_id" json:"showtime_id"`
	StatusTicket string    `db:"status_ticket" json:"status_ticket"` // "active", "not_active"
	StatusPaid   string    `db:"status_paid" json:"status_paid"`     // "paid" "not_paid"
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

type OrderHistoryDetail struct {
	Booking  Booking
	Showtime Showtime
	Movie    Movie
	Cinema   Cinema
}
