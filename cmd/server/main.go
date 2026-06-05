package main

import (
	"log"

	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/api"
	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/app"
	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/dataset"
	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/fraud"
	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/worker"
)

func main() {
	Init()

	mux := api.NewMux()

	for i := 0; i < 2; i++ {
		go worker.Run("/tmp/lb.sock", mux)
	}

	select {}
}

func Init() {
	ds, err := dataset.Load(
		"resources/references.json.gz",
		"resources/normalization.json",
		"resources/mcc_risk.json",
	)
	if err != nil {
		log.Fatal(err)
	}

	app.Vectorize = fraud.NewVectorizer(ds.Normalization, ds.MccRisk)
	index := fraud.NewBruteForce(ds.References)
	app.Detector = fraud.NewDetector(index)

	app.Ready.Store(true)

	log.Println("application ready")
}
