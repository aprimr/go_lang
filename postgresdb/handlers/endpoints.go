package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/aprimr/goanddatabase/db"
	"github.com/aprimr/goanddatabase/models"
	"github.com/aprimr/goanddatabase/utils"
)

func InsertTodoHandler(w http.ResponseWriter, r *http.Request, database *sql.DB) {
	// POST Method => "/todos" : insert todo into db
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

func FetchAllTodosHandler(w http.ResponseWriter, r *http.Request, database *sql.DB) {
	// GET Method => "/todos" : fetch all todos from db
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

func FetchTodosByIDHandler(w http.ResponseWriter, r *http.Request, database *sql.DB) {
	// GET Method => "/todos/:id" : fetch todos by id
	// Check if request method matches
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", 500)
		return
	}

	// Parse url
	urlStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(urlStr)
	if err != nil {
		panic(err)
	}

	// Fetch and validate todo
	todo, err := db.FetchTodosByID(database, id)
	if err != nil {
		panic(err)
	}
	if todo == nil {
		utils.EncodeJson(w, map[string]any{
			"message": "Todo fetch unsuccessful",
			"success": false,
			"data":    nil,
		}, 404)
		return
	}

	utils.EncodeJson(w, map[string]any{
		"message": "Todo fetch successful",
		"success": true,
		"data":    todo,
	}, 200)
}

func DeleteTodosByIDHandler(w http.ResponseWriter, r *http.Request, database *sql.DB) {
	// DELETE Method => "/todos/:id" : delete todos by id
	// Check if request method matches
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", 500)
		return
	}

	// Parse url
	urlStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(urlStr)
	if err != nil {
		panic(err)
	}

	err = db.DeleteTodosByID(database, id)
	if err != nil {
		// return error json
		utils.EncodeJson(w, map[string]any{
			"message": "Error deleting todo",
			"success": false,
		}, 404)
		return
	}

	utils.EncodeJson(w, map[string]any{
		"message": "Todo deleted successfully",
		"success": true,
	}, 200)
}
