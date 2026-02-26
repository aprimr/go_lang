// TODO APP

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// struct for todo
type Todo struct {
	Id          int    `json:"todoId"`
	Title       string `json:"todoTitle"`
	IsCompleted bool   `json:"isCompleted"`
}

var todos []Todo
var nextId = 1

func createTodo(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is not POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	var todo Todo

	// Decode the JSON object from request body
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Error decoding JOSN", http.StatusBadRequest)
		return
	}

	todo.Id = nextId
	todo.IsCompleted = false
	nextId++

	todos = append(todos, todo)

	// create response return
	resReturn := map[string]any{
		"message": "Todo created successfully",
		"success": true,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resReturn)
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is not GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	// create response return
	resReturn := map[string]any{
		"message": "Todo fetched successfully",
		"todos":   todos,
		"success": true,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resReturn)
}

func main() {
	port := ":8000"
	mux := http.NewServeMux()

	// Create todo api - POST method
	mux.HandleFunc("/todo", createTodo)

	// Get all todos api - GET method
	mux.HandleFunc("/todos", getTodos)

	fmt.Println("Server running on port", port)
	http.ListenAndServe(port, mux)

}
