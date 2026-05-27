package serverContext

import (
	"encoding/json"
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

func ValidatePayload[T any](r *http.Request) (payload T, err error) {
	err = json.NewDecoder(r.Body).Decode(&payload)
	return
}
