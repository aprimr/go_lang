package main

import (
	"chi-router/handler"
	"chi-router/middleware"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	// create chi router
	r := chi.NewRouter()
	// r.Use(middleware.Logger)
	r.Use(chimiddleware.Recoverer)

	// Custom middleware
	r.Use(middleware.Logger)

	// Routes
	// r.Get("/home", handler.HandlerHomeRoute)
	// r.Get("/user/{id}", handler.UserHandler)
	// r.Get("/user/{id}/chats/{chatid}", handler.HandleUserChat)

	// route groups
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", handler.HandlerHomeRoute)
			r.Get("/{id}", handler.UserHandler)
		})

		r.Route("/post", func(r chi.Router) {
			r.Get("/{id}", handler.GetPostById)
			r.Post("/", handler.CeateAPost)
		})
	})

	// Start server
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
