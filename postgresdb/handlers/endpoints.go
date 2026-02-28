package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/aprimr/goanddatabase/db"
	"github.com/aprimr/goanddatabase/models"
	"github.com/aprimr/goanddatabase/utils"
)

func InsertTodoHandler(w http.ResponseWriter, r *http.Request, database *sql.DB) {
	// POST Method => "/todo" : insert todo into db
	// Check if request method matches
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", 500)
		return
	}

	// Decode JSON
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Invalid JSON", 500)
		return
	}

	// Insert Into DB
	db.InsertTodo(database, todo.Title, todo.Is_completed)
	utils.EncodeJson(w, map[string]any{
		"message": "Todo created",
		"success": true,
	}, 201)
}

func FetchAllTodos(w http.ResponseWriter, r *http.Request, database *sql.DB) {
	// GET Method => "/todo" : fetch all todos from db
	// Check if request method matches
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", 500)
		return
	}

	todos, err := db.FetchAllTodos(database)
	if err != nil {
		panic(err)
	}

	utils.EncodeJson(w, map[string]any{
		"message": "Todos fetch successful",
		"success": true,
		"data":    todos,
	}, 200)
}
