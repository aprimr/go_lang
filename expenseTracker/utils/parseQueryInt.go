package utils

import (
	"net/http"
	"strconv"
)

func ParseQueryInt(r *http.Request, key string, defaultValue int) int {
	valStr := r.URL.Query().Get(key)
	if valStr == "" {
		return defaultValue
	}

	// convert val to int
	valInt, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultValue
	}

	return valInt
}
