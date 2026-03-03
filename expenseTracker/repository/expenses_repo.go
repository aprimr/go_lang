package repository

import (
	"context"
	"expenseTracker/models"
	"fmt"
	"math"

	"github.com/jackc/pgx/v5"
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
		Limit:      limit,
		TotalCount: totalRows,
		TotalPages: totalPages,
	}, nil
}

// Get Expense by id
func GetExpenseById(db *pgxpool.Pool, id int) (models.Expense, error) {
	expense := models.Expense{}

	// Fire query
	row := db.QueryRow(context.Background(), "SELECT id, title, amount, category, date, created_at FROM expenses WHERE id=$1", id)
	err := row.Scan(&expense.Id, &expense.Title, &expense.Amount, &expense.Category, &expense.Date, &expense.CreatedAt)
	if err == pgx.ErrNoRows {
		return models.Expense{}, fmt.Errorf("No expenses found")
	}
	if err != nil {
		return models.Expense{}, err
	}

	return expense, nil
}

// Update Expense
func UpdateExpense(db *pgxpool.Pool, id int, expense models.Expense) error {
	commandTag, err := db.Exec(context.Background(), "UPDATE expenses SET title=$1, amount=$2, category=$3, date=$4 WHERE id=$5", expense.Title, expense.Amount, expense.Category, expense.Date, id)

	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("Error updating expense")
	}

	return nil
}

// Delete Expense
func DeleteExpense(db *pgxpool.Pool, id int) error {
	commandTag, err := db.Exec(context.Background(), "DELETE FROM expenses WHERE id=$1", id)

	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("Error deleting expense")
	}

	return nil
}

// Get Expense Summary
func GetExpenseSummary(db *pgxpool.Pool) (models.ExpenseSummary, error) {
	var expenseSummary models.ExpenseSummary
	// get total amount
	err := db.QueryRow(context.Background(), "SELECT COALESCE(SUM(amount), 0) FROM expenses").Scan(&expenseSummary.TotalSpent)
	if err != nil {
		return models.ExpenseSummary{}, err
	}

	// get category and its total
	rows, err := db.Query(context.Background(), "SELECT category, SUM(amount) AS total FROM expenses GROUP BY category ORDER BY total DESC")
	if err != nil {
		return models.ExpenseSummary{}, err
	}

	// Scan each row
	var categorySummary []models.CategorySummary
	for rows.Next() {
		var category models.CategorySummary

		err = rows.Scan(&category.Category, &category.Total)
		if err != nil {
			return models.ExpenseSummary{}, err
		}

		// append category into slice of categories
		categorySummary = append(categorySummary, category)
	}

	expenseSummary.ByCategory = categorySummary
	return expenseSummary, err
}
