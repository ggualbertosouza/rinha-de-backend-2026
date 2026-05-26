package server

import (
	"net/http"
)

func ReadyHandler(w http.ResponseWriter, r *http.Request) {
	if !AllowMethod(w, r, http.MethodGet) {
		return
	}

	WriteResponse(w, "Ready handler")
}

func FraudScoreHandler(w http.ResponseWriter, r *http.Request) {
	if !AllowMethod(w, r, http.MethodGet) {
		return
	}

	SetJSON(w)

	WriteResponse(w, "Fraud score handler")
}
