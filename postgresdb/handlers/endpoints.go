package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/aprimr/goanddatabase/db"
	"github.com/aprimr/goanddatabase/models"
	"github.com/aprimr/goanddatabase/utils"
)

func InsertTodoHandler(w http.ResponseWriter, r *http.Request, database *sql.DB) {
	// POST Method => "/todos" : insert todo into db

	// Decode JSON
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	// Insert Into DB
	err = db.InsertTodo(database, todo.Title, todo.Is_completed)
	if err != nil {
		utils.EncodeJson(w, map[string]any{
			"message": "Database error",
			"success": false,
		}, 500)
		return
	}
	utils.EncodeJson(w, map[string]any{
		"message": "Todo created",
		"success": true,
	}, 201)
}

func FetchAllTodosHandler(w http.ResponseWriter, r *http.Request, database *sql.DB) {
	// GET Method => "/todos" : fetch all todos from db

	todos, err := db.FetchAllTodos(database)
	if err != nil {
		http.Error(w, "Cannot parse url", http.StatusBadRequest)
		return
	}

	utils.EncodeJson(w, map[string]any{
		"message": "Todos fetch successful",
		"success": true,
		"data":    todos,
	}, 200)
}

func FetchTodosByIDHandler(w http.ResponseWriter, r *http.Request, database *sql.DB) {
	// GET Method => "/todos/:id" : fetch todos by id

	// Parse url
	urlStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(urlStr)
	if err != nil {
		http.Error(w, "Cannot parse url", http.StatusBadRequest)
		return
	}

	// Fetch and validate todo
	todo, err := db.FetchTodosByID(database, id)
	if err != nil {
		log.Printf("FetchTodoByID: db error: %v", err)
		utils.EncodeJson(w, map[string]any{
			"message": "Failed to fetch todo",
			"success": false,
		}, http.StatusInternalServerError)
		return
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

	// Parse url
	urlStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(urlStr)
	if err != nil {
		http.Error(w, "Cannot parse url", http.StatusBadRequest)
		return
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

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request, database *sql.DB) {
	// PUT Method => "/todos/:id" : update todos

	// Parse url
	urlStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(urlStr)
	if err != nil {
		http.Error(w, "Cannot parse url", http.StatusBadRequest)
		return
	}

	// Parse JSON
	var todo models.Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err = db.UpdateTodo(database, id, todo)
	if err != nil {
		// return error json
		utils.EncodeJson(w, map[string]any{
			"message": "Error updating todo",
			"success": false,
		}, 500)
		return
	}

	utils.EncodeJson(w, map[string]any{
		"data":    todo,
		"message": "Todo updated",
		"success": true,
	}, 200)
}
