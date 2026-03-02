package utils

import (
	"encoding/json"
	"expenseTracker/models"
	"net/http"
)

func SendSuccess(w http.ResponseWriter, message string, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(models.SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func SendError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Success: true,
		Message: message,
	})
}
