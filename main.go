package main

import (
	"log"

	server "github.com/ggualbertosouza/rinha-de-backend-2026/internal/http"
)

func main() {
	srv := server.New("8080")

	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
