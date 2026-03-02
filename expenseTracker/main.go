package main

import (
	"expenseTracker/db"
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
	_, err = db.ConnectDB()
	if err != nil {
		log.Panic("Error connecting to db. ", err)
	}

	// Create new serveMux
	mux := http.NewServeMux()

	// SpinUp server
	serverPort := ":" + os.Getenv("SERVER_PORT")
	log.Println("Server running on port", serverPort)
	err = http.ListenAndServe(serverPort, mux)
	if err != nil {
		log.Panic("Error running server. ", err)
	}
}
