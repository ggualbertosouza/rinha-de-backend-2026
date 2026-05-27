package server

import (
	"fmt"
	"net/http"

	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/app"
)

func ReadyHandler(w http.ResponseWriter, r *http.Request) {
	if !AllowMethod(w, r, http.MethodGet) {
		return
	}

	if !app.Ready.Load() {
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func FraudScoreHandler(w http.ResponseWriter, r *http.Request) {
	if !AllowMethod(w, r, http.MethodGet) {
		return
	}

	SetJSON(w)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Fraud score handler")
}
