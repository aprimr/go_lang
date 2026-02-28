package main

import (
	"fmt"
	"net/http"

	"github.com/aprimr/goanddatabase/db"
	"github.com/aprimr/goanddatabase/handlers"

	_ "github.com/lib/pq"
)

func main() {
	port := ":8099"
	mux := http.NewServeMux()

	// Connect to database
	database, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer database.Close()

	// API Endpoints
	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// POST Method -> insert todo into db
		case http.MethodPost:
			handlers.InsertTodoHandler(w, r, database)

		// GET Method -> fetch all todos from db
		case http.MethodGet:
			handlers.FetchAllTodosHandler(w, r, database)

		// handle default case
		default:
			http.Error(w, "Method not allowed", 500)
		}
	})

	mux.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// GET Method -> fetch todos by id
		case http.MethodGet:
			handlers.FetchTodosByIDHandler(w, r, database)
		}
	})

	// SpinUp the server
	fmt.Println("Server started on port", port)
	err = http.ListenAndServe(port, mux)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
