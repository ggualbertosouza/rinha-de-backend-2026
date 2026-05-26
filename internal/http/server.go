package server

import (
	"log"
	"net/http"
	"time"
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

	mux.HandleFunc("/ready", ReadyHandler)
	mux.HandleFunc("/fraud-score", FraudScoreHandler)

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
