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
	Data       []Expense `json:"data"`
	Page       int       `json:"page"`
	Limit      int       `json:"limit"`
	TotalCount int       `json:"total_count"`
	TotalPages int       `json:"total_pages"`
}

type CategorySummary struct {
	Category string  `json:"category"`
	Total    float64 `json:"total"`
}

type ExpenseSummary struct {
	TotalSpent float64           `json:"total_spent"`
	ByCategory []CategorySummary `json:"by_category"`
}
