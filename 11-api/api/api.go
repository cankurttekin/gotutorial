package api

import (
	"encoding/json"
	"net/http"
)

// coin balance params
type CoinBalanceParams struct {
	Username string
}

// coin balance response
type CoinBalanceResponse struct {
	// success
	Code int

	// account balance
	Balance int64
}

// error response
type Error struct {
	// error code
	Code int
	
	// error message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	response := Error{
		Code:    code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(response)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An unexpected error occurred", http.StatusInternalServerError)
	}
)









		





