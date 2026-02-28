package main

import (
	"fmt"
	"net/http"

	"github.com/aprimr/goanddatabase/db"

	_ "github.com/lib/pq"
)

func main() {
	port := ":8000"
	mux := http.NewServeMux()

	// Connect to database
	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// SpinUp the server
	fmt.Println("Server started on port", port)
	err = http.ListenAndServe(port, mux)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
