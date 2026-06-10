package app

import (
	"log"

	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/dataset"
	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/fraud"
)

func Init() {
	ds, err := dataset.Load(
		"/resources/references.json.gz",
		"/resources/normalization.json",
		"/resources/mcc_risk.json",
	)
	if err != nil {
		log.Fatal(err)
	}

	Vectorize = fraud.NewVectorizer(ds.Normalization, ds.MccRisk)
	index := fraud.NewBruteForce(ds.References)
	Detector = fraud.NewDetector(index)

	Ready.Store(true)

	log.Println("application ready")
}
