package apiContext

import (
	"encoding/json"
	"net/http"
)

func SetJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func AllowMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func ValidatePayload[T any](r *http.Request) (T, error) {
	var payload T
	err := json.NewDecoder(r.Body).Decode(&payload)
	return payload, err
}
