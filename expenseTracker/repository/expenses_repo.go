package repository

import (
	"context"
	"expenseTracker/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Add Expenses
func AddExpenses(db *pgxpool.Pool, expense models.Expense) error {
	_, err := db.Exec(context.Background(), "INSERT INTO expenses (title, amount, category, date) VALUES($1, $2, $3, $4)", expense.Title, expense.Amount, expense.Category, expense.Date)
	return err
}
