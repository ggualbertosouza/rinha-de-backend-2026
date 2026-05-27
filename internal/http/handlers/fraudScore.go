package handlers

import (
	"fmt"
	"net/http"

	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/app"
	serverContext "github.com/ggualbertosouza/rinha-de-backend-2026/internal/http/context"
)

func FraudScoreHandler(w http.ResponseWriter, r *http.Request) {
	if !serverContext.AllowMethod(w, r, http.MethodPost) {
		return
	}
	serverContext.SetJSON(w)

	_, err := serverContext.ValidatePayload[[]app.InputPayload](r)
	if err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	/*
		1. Validar http method 				[x]
		2. Validar payload - json 			[x]
		3. Transformar paylaod em vetor		[ ]
		4. Busca vetorial 					[ ]
		5. Calcular score					[ ]
		6. Responder 						[ ]
	*/

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Fraud score handler")
}
