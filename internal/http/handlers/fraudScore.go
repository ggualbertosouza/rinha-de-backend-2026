package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/app"
	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/fraud"
	serverContext "github.com/ggualbertosouza/rinha-de-backend-2026/internal/http/context"
)

func FraudScoreHandler(w http.ResponseWriter, r *http.Request) {
	if !app.Ready.Load() {
		return
	}

	if !serverContext.AllowMethod(w, r, http.MethodPost) {
		return
	}
	serverContext.SetJSON(w)

	payload, err := serverContext.ValidatePayload[fraud.Payload](r)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	query := app.Vectorize.Vectorize(payload)

	result := app.Detector.Detect(query)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
