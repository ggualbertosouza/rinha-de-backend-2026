package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/app"
	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/dataset"
	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/fraud"
	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/http/handlers"
)

func main() {
	setup()

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
	body, _ := json.Marshal(req)

	r := httptest.NewRequest(http.MethodPost, "/fraud-score", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handlers.FraudScoreHandler(rr, r)

	log.Printf(
		"id=%s status=%d body=%q",
		req.ID,
		rr.Code,
		rr.Body.String(),
	)
}

func setup() {
	ds, err := dataset.Load(
		"resources/references.json.gz",
		"resources/normalization.json",
		"resources/mcc_risk.json",
	)
	if err != nil {
		log.Fatal(err)
	}

	app.Ready.Store(true)

	app.Vectorize = fraud.NewVectorizer(
		ds.Normalization,
		ds.MccRisk,
	)

	index := fraud.NewBruteForce(
		ds.References,
	)

	app.Detector = fraud.NewDetector(index)
}
