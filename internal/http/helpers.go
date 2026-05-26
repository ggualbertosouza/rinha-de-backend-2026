package server

import (
	"fmt"
	"net/http"
)

func SetJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func AllowMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}

	return true
}

func WriteResponse(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, message)
}
