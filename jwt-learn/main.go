package main

import (
	"learn-jwt/handlers"
	"learn-jwt/middleware"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Env file not loaded")
	}

	// Create chi router
	r := chi.NewRouter()

	// Create Routes
	r.Route("/jwt/v1", func(r chi.Router) {
		r.Post("/login", handlers.LoginHandler)
		r.Post("/register", handlers.RegisterHandler)

		r.Group(func(r chi.Router) {
			r.Use(middleware.Authorization)
			r.Get("/profile", handlers.FetchProfileById)
		})

	})
	// Start Server
	port := ":" + os.Getenv("PORT")
	log.Println("Server started on port", port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		log.Fatalln("Error starting server")
	}
}
