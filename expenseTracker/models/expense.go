package models

import "time"

type Expense struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Amount    float64   `json:"amount"`
	Category  string    `json:"category"`
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
}

type PaginatedExpenses struct {
	Data       any `json:"data"`
	Page       int `json:"page"`
	TotalCount int `json:"total_count"`
	TotalPages int `json:"total_pages"`
}
