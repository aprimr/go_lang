package models

type Expense struct {
	Id        int     `json:"id"`
	Title     string  `json:"title"`
	Amount    float64 `json:"amount"`
	Category  string  `json:"category"`
	Date      string  `json:"date"`
	CreatedAt string  `json:"created_at"`
}
