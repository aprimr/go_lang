package handlers

import (
	"encoding/json"
	"expenseTracker/models"
	"expenseTracker/repository"
	"expenseTracker/utils"
	"log"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

// POST -> /expenses
func AddExpensesHandler(w http.ResponseWriter, r *http.Request, db *pgxpool.Pool) {
	// Parse JSON (title, amount, category, date)
	var expense models.Expense
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		log.Printf("Invalid JSON: %v", err)
		utils.SendError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// validate the data
	if strings.TrimSpace(expense.Title) == "" {
		utils.SendError(w, "Title is required", http.StatusBadRequest)
		return
	}
	if expense.Amount < 0 {
		expense.Amount = 0
	}
	if strings.TrimSpace(expense.Category) == "" {
		utils.SendError(w, "Category is required", http.StatusBadRequest)
		return
	}

	// Call AddExepnses
	if err := repository.AddExpenses(db, expense); err != nil {
		log.Printf("AddExpenses -> db error: %v", err)
		utils.SendError(w, "Failed to create expense", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, "Expense created", nil, http.StatusCreated)
}

// GET -> /expenses?page=1&limit=10
func GetExpensesHandler(w http.ResponseWriter, r *http.Request, db *pgxpool.Pool) {
	//  get queryParams
	page := utils.ParseQueryInt(r, "page", 1)
	limit := utils.ParseQueryInt(r, "limit", 10)

	// Validate Query Params
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	// Call GetExpenses
	expenses, err := repository.GetExpenses(db, page, limit)
	if err != nil {
		log.Printf("GetExpenses -> db error: %v", err)
		utils.SendError(w, "Error fetching expenses", http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, "Expenses fetched successfully", expenses, http.StatusOK)
}
