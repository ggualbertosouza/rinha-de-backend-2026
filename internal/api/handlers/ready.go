package handlers

import (
	"net/http"

	serverContext "github.com/ggualbertosouza/rinha-de-backend-2026/internal/api/context"
	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/app"
)

func ReadyHandler(w http.ResponseWriter, r *http.Request) {
	if !serverContext.AllowMethod(w, r, http.MethodGet) {
		return
	}

	serverContext.SetJSON(w)

	if !app.Ready.Load() {
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte(`{"status":"fail"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}
