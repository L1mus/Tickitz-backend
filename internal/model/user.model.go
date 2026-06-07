package model

import "time"

type Users struct {
	ID          int         `json:"id" db:"id"`
	Email       string      `json:"email" db:"email"`
	Password    string      `json:"-" db:"password"`
	First_Name  string      `json:"first_name" db:"first_name"`
	Last_Name   string      `json:"last_name" db:"last_name"`
	Phone       *string      `json:"phone" db:"phone"`
	Photo       *string      `json:"photo" db:"photo"`
	Role        *string      `json:"role" db:"role"`
	Location    *Locations  `json:"location" db:"location_id"`
	Is_Active   bool        `json:"is_active" db:"is_active"`
	Created_At  time.Time    `json:"created_at" db:"created_at"`
	Updated_At  time.Time    `json:"updated_at" db:"updated_at"`
}

type Locations struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

