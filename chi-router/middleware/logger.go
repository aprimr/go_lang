package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("%v: New Request -> %v, %v \n", time.DateTime, r.URL.Path, r.Method)
		next.ServeHTTP(w, r)
		fmt.Println("Time taken:", time.Since(start))
	})
}
