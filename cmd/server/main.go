package main

import (
	"log"

	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/app"
	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/dataset"
	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/fraud"
	server "github.com/ggualbertosouza/rinha-de-backend-2026/internal/http"
)

func main() {
	ds, err := dataset.Load(
		"resources/references.json.gz",
		"resources/normalization.json",
		"resources/mcc_risk.json",
	)
	if err != nil {
		log.Fatal(err)
	}

	app.Ready.Store(true)

	app.Vectorize = fraud.NewVectorizer(ds.Normalization, ds.MccRisk)
	index := fraud.NewBruteForce(ds.References)
	app.Detector = fraud.NewDetector(index)

	srv := server.New("9999")

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}

/*
	json basic
	2026/06/04 11:23:59 loaded 3000000 references in 6.894391495s
	2026/06/04 11:23:59 load normalization in 47.11µs
	2026/06/04 11:23:59 load mccRisk in 14.8µs
	2026/06/04 11:23:59 server starting on 9999

	Sonic
	
*/