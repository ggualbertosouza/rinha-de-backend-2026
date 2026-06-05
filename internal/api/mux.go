package api

import (
	"net/http"

	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/api/handlers"
)

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/ready", handlers.ReadyHandler)
	mux.HandleFunc("/fraud-score", handlers.FraudScoreHandler)

	return mux
}
