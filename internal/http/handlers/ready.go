package handlers

import (
	"net/http"

	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/app"
	serverContext "github.com/ggualbertosouza/rinha-de-backend-2026/internal/http/context"
)

func ReadyHandler(w http.ResponseWriter, r *http.Request) {
	if !serverContext.AllowMethod(w, r, http.MethodGet) {
		return
	}
	serverContext.SetJSON(w)

	if !app.Ready.Load() {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"status":"fail"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
