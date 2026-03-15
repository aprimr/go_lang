package middleware

import (
	"context"
	"encoding/json"
	"learn-jwt/utils"
	"net/http"
	"strings"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get auth header and trim bearer from jwt
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("missing token")
			return
		}
		jwtToken := strings.TrimPrefix(authHeader, "Bearer ")

		// Verify jwt token
		payload, err := utils.VerifyToken(jwtToken)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Jwt token tampered")
			return
		}

		// call the handler if everything is ok
		ctx := context.WithValue(r.Context(), "user_id", payload.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
