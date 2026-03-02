package handlers

import (
	"encoding/json"
	"expenseTracker/models"
	"expenseTracker/repository"
	"expenseTracker/utils"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

// POST -> /expenses
func AddExpensesHandler(w http.ResponseWriter, r *http.Request, db *pgxpool.Pool) {
	// Parse JSON (title, amount, category, date)
	var expesne models.Expense
	err := json.NewDecoder(r.Body).Decode(&expesne)
	if err != nil {
		utils.SendError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// validate the data
	if strings.TrimSpace(expesne.Title) == "" {
		utils.SendError(w, "Title is required", http.StatusBadRequest)
		return
	}
	if expesne.Amount < 0 {
		expesne.Amount = 0
	}
	if strings.TrimSpace(expesne.Category) == "" {
		utils.SendError(w, "Category is required", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(expesne.Date) == "" {
		utils.SendError(w, "Date is required", http.StatusBadRequest)
		return
	}

	// Call AddExepnses
	if err := repository.AddExpenses(db, expesne); err != nil {
		utils.SendError(w, "Failed to create expense", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, "Expense created", nil, http.StatusCreated)
}
