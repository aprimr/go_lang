package handlers

import (
	"encoding/json"
	"learn-jwt/store"
	"net/http"
)

func FetchProfileById(w http.ResponseWriter, r *http.Request) {
	// Get id from URl
	id := r.Context().Value("user_id").(string)

	// Fetch user details by id
	user, exist := store.GetUserByid(id)
	if exist == false {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
