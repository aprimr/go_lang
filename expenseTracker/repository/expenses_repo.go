package repository

import (
	"context"
	"expenseTracker/models"
	"math"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Add Expenses
func AddExpenses(db *pgxpool.Pool, expense models.Expense) error {
	_, err := db.Exec(context.Background(), "INSERT INTO expenses (title, amount, category, date) VALUES($1, $2, $3, $4)", expense.Title, expense.Amount, expense.Category, expense.Date)
	return err
}

// Get All Expenses
func GetExpenses(db *pgxpool.Pool, page int, limit int) (models.PaginatedExpenses, error) {
	// Calculate offset value
	offset := (page - 1) * limit

	// Get no of rows
	var totalRows int
	err := db.QueryRow(context.Background(), "SELECT COUNT(*) FROM expenses").Scan(&totalRows)
	if err != nil {
		return models.PaginatedExpenses{}, err
	}

	// Fire Query
	rows, err := db.Query(context.Background(), "SELECT id, title, amount, category, date, created_at FROM expenses ORDER BY id LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return models.PaginatedExpenses{}, err
	}
	defer rows.Close()

	// Loop through rows and scan each result
	expenses := []models.Expense{}

	for rows.Next() {
		var expense models.Expense

		err := rows.Scan(&expense.Id, &expense.Title, &expense.Amount, &expense.Category, &expense.Date, &expense.CreatedAt)
		if err != nil {
			return models.PaginatedExpenses{}, err

		}

		expenses = append(expenses, expense)
	}

	// Calculate total Pages
	totalPages := int(math.Ceil(float64(totalRows) / float64(limit)))

	return models.PaginatedExpenses{
		Data:       expenses,
		Page:       page,
		TotalCount: totalRows,
		TotalPages: totalPages,
	}, nil
}
