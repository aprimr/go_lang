package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type URLRequest struct {
	Url string `json:"url"`
}

var urlStore = make(map[string]string)

func generateShortId(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}

func createErrorMesage(w http.ResponseWriter, message string) {
	// create error message
	errorRes := map[string]any{
		"message": message,
	}

	// send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(errorRes)
}

func createSuccessMesage(w http.ResponseWriter, message string, body any) {
	// create success message
	successRes := map[string]any{
		"message": message,
		"body":    body,
	}

	// send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(successRes)
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	// check if the request method is not POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	// Parse JSON from request body
	var urlReq URLRequest
	err := json.NewDecoder(r.Body).Decode(&urlReq)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	r.Body.Close()

	// Validate URL
	if strings.TrimSpace(urlReq.Url) == "" {
		createErrorMesage(w, "Please enter a URL")
		return
	}

	_, ValidateErr := url.ParseRequestURI(urlReq.Url)
	if ValidateErr != nil {
		createErrorMesage(w, "Please enter a valid URL")
		return
	}

	// Create new map
	shortId := generateShortId(6)
	urlStore[shortId] = urlReq.Url
	bodyMsg := fmt.Sprintf("Short Url: http://localhost:8000/%s", shortId)
	createSuccessMesage(w, "URL shortened successfully", bodyMsg)
}

func main() {
	port := ":8080"
	mux := http.NewServeMux()

	// API endpoint for shorten
	mux.HandleFunc("/shorten", shortenHandler)

	fmt.Println("Server started on port", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
