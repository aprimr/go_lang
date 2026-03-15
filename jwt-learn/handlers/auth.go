package handlers

import (
	"encoding/json"
	"learn-jwt/models"
	"learn-jwt/store"
	"learn-jwt/utils"
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Parse Request body
	registerData := models.RegisterDetails{}
	err := json.NewDecoder(r.Body).Decode(&registerData)
	log.Println(registerData.Password)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ERROR DECODING JSON")
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerData.Password), bcrypt.DefaultCost)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	// Create user var
	user := models.User{}

	// Create an unique user id
	user.Id = uuid.New().String()
	user.Name = registerData.Name
	user.Email = registerData.Email

	// Replace the user password with hash
	user.Password = string(hashedPassword)

	// Add user in user store
	store.AddUser(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("User registered successfully")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Parse Request Body
	loginDetails := models.LoginDetails{}
	err := json.NewDecoder(r.Body).Decode(&loginDetails)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("ERROR DECODING JSON")
		return
	}

	// Find user by email
	user, exists := store.GetUserByEmail(loginDetails.Email)
	if exists == false {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("User not found")
		return
	}

	// Compare the password in store with login details password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDetails.Password))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Incorrect password")
		return
	}

	// Create a JWT token
	jwtToken, err := utils.CreateToken(user.Id, "role")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Something went wrong")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(jwtToken)
}
