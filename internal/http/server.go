package server

import (
	"log"
	"net/http"
	"time"

	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/http/handlers"
)

type Server struct {
	Port string
}

func New(port string) *Server {
	return &Server{
		Port: port,
	}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/ready", handlers.ReadyHandler)
	mux.HandleFunc("/fraud-score", handlers.FraudScoreHandler)

	srv := &http.Server{
		Addr:              ":" + s.Port,
		Handler:           mux,
		ReadTimeout:       3 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       30 * time.Second,
		MaxHeaderBytes:    1 << 10,
	}

	log.Printf("server starting on %s", s.Port)
	return srv.ListenAndServe()
}
