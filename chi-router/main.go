package main

import (
	"chi-router/handler"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	// create chi router
	r := chi.NewRouter()

	// Routes
	r.Get("/home", handler.HandlerHomeRoute)
	r.Get("/user/{id}", handler.UserHandler)
	r.Get("/user/{id}/chats/{chatid}", handler.HandleUserChat)

	// Start server
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
