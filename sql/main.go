package main

import (
	"log"
	"net/http"
	"os"
	"sql/db"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env files")
	}
	db.ConnectDB()
	port := "." + os.Getenv("PORT")
	mux := http.NewServeMux()

	//

	log.Println("Server started on port", port)
	http.ListenAndServe(port, mux)
}
