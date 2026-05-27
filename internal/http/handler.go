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
	SetJSON(w)

	if !app.Ready.Load() {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"status":"fail"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

func FraudScoreHandler(w http.ResponseWriter, r *http.Request) {
	if !AllowMethod(w, r, http.MethodPost) {
		return
	}
	SetJSON(w)

	/*
		1. Validar http method 				[x]
		2. Validar payload - json 			[x]
		3. Transformar paylaod em vetor		[x]
		4. Busca vetorial 					[x]
		5. Calcular score					[x]
		6. Responder 						[x]
	*/

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Fraud score handler")
}
