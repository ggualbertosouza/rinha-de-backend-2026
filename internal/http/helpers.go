package server

import (
	"net/http"
)

var (
	headerJSON = []string{"application/json"}
)

func SetJSON(w http.ResponseWriter) {
	w.Header()["Content-Type"] = headerJSON
}

func AllowMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}

	return true
}
