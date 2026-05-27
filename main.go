package main

import (
	"log"

	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/app"
	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/dataset"
	server "github.com/ggualbertosouza/rinha-de-backend-2026/internal/http"
)

func main() {
	go loadDataset()

	srv := server.New("8080")

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}

func loadDataset() {
	log.Print("processing dataset")

	references, err := dataset.Load("resources/references.json.gz")
	if err != nil {
		log.Fatal(err)
	}

	app.Ready.Store(true)
	log.Printf("dataset loaded: %d references", len(references))
}
