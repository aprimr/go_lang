// TODO APP

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
	r.Body.Close()

	// Validate data
	if strings.TrimSpace(todo.Title) == "" {
		http.Error(w, "Title cannot be empty", http.StatusBadRequest)
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

	w.WriteHeader(http.StatusOK)
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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resReturn)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is not GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	// parse the path
	idStr := strings.TrimPrefix(r.URL.Path, "/todo/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	for _, todo := range todos {
		if todo.Id == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(todo)
			return
		}
	}

	http.Error(w, "Id not found.", http.StatusNotFound)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	// Check if the request mrthod is not DELETE
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	// parse path
	idStr := strings.TrimPrefix(r.URL.Path, "/todo/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[:i], todos[i+1:]...)

			// create response return
			resReturn := map[string]any{
				"message": "Todo deleted successfully",
				"todos":   todos,
				"success": true,
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resReturn)
			return
		}
	}
	http.Error(w, "Id not matched", http.StatusNotFound)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is not PUT
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	var todo Todo

	// Parse JSON
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}
	r.Body.Close()

	// Parse url
	idStr := strings.TrimPrefix(r.URL.Path, "/todo/")
	id, errr := strconv.Atoi(idStr)
	if errr != nil {
		http.Error(w, "Invalid Id", http.StatusBadRequest)
		return
	}

	for i, t := range todos {
		if t.Id == id {
			todos[i].Title = todo.Title
			todos[i].IsCompleted = todo.IsCompleted

			resReturn := map[string]any{
				"message": "Todo Updated successfully",
				"todos":   todos,
				"success": true,
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resReturn)
			return
		}
	}

	http.Error(w, "Id not found", http.StatusNotFound)
}

func main() {
	port := ":8000"
	mux := http.NewServeMux()

	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			// Create todo route - POST method
			createTodo(w, r)

		case http.MethodGet:
			// Get all todos route - GET method
			getTodos(w, r)
		}
	})

	mux.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// Get todo by id route - GET method
		case http.MethodGet:
			getTodo(w, r)

		// Delete todo by id route - DELETE method
		case http.MethodDelete:
			deleteTodo(w, r)

		// Update todo by id route - PUT method
		case http.MethodPut:
			updateTodo(w, r)
		}
	})

	fmt.Println("Server running on port", port)
	http.ListenAndServe(port, mux)

}
