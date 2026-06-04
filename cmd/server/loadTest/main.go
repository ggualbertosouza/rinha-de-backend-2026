package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/fraud"
	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/http/handlers"
)

func main() {
	payloads, err := loadPayloads("resources/example-payloads.json")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("requests length: %d", len(payloads))

	for _, p := range payloads {
		executeRequest(p)
	}
}

func loadPayloads(path string) ([]fraud.Payload, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var payloads []fraud.Payload

	err = json.Unmarshal(data, &payloads)
	if err != nil {
		return nil, err
	}

	return payloads, nil
}

func executeRequest(req fraud.Payload) {
	body, _ := json.Marshal([]fraud.Payload{req})

	r := httptest.NewRequest(http.MethodPost, "/fraud-score", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handlers.FraudScoreHandler(rr, r)

	log.Printf(
		"id=%s status=%d body=%v",
		req.ID,
		rr.Code,
		rr.Body,
	)
}
