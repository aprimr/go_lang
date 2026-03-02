package main

import (
	"expenseTracker/db"
	"expenseTracker/handlers"
	"expenseTracker/utils"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading env")
	}

	// Connect to db
	database, err := db.ConnectDB()
	if err != nil {
		log.Panic("Error connecting to db. ", err)
	}

	// Create new serveMux
	mux := http.NewServeMux()

	// Endpoints
	mux.HandleFunc("/expenses", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// POST METHOD (`/expenses`) -> add expenses
		case http.MethodPost:
			handlers.AddExpensesHandler(w, r, database)

		default:
			utils.SendError(w, "Method not allowed", http.StatusBadRequest)
		}
	})

	// SpinUp server
	serverPort := ":" + os.Getenv("SERVER_PORT")
	log.Println("Server running on port", serverPort)
	err = http.ListenAndServe(serverPort, mux)
	if err != nil {
		log.Panic("Error running server. ", err)
	}
}
