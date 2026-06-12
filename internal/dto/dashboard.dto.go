package dto

type SalesChartFilter struct {
	MovieName string `form:"movie_name"`
	FilterBy  string `form:"filter_by" binding:"required,oneof=weekly monthly"`
}

type SalesChartResponse struct {
	Label        string  `json:"label"`
	TotalRevenue float64 `json:"total_revenue"`
}

type TicketSalesFilter struct {
	GenreID    int `form:"genre_id"`
	LocationID int `form:"location_id"`
}

type TicketSalesResponse struct {
	MovieTitle string `json:"movie_title"`
	TotalSold  int    `json:"total_sold"`
}
