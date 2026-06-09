package model

import "time"

type SeatRow struct {
	SeatID     int
	Row        string
	SeatNumber int
	SeatType   string
	SeatStatus string
}

type ShowtimeSummary struct {
	MovieTitle  string
	MoviePoster string
	Category    string
	CinemaID    int
	CinemaName  string
	CinemaLogo  string
	ShowDate    time.Time
	ShowTime    string
	TicketPrice int
}
