package main

import (
	"os"
	"strconv"

	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/api"
	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/api/middleware"
	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/app"
	"github.com/ggualbertosouza/rinha-de-backend-2026/internal/worker"
)

func main() {
	sp := os.Getenv("LB_SOCKET")
	wa := os.Getenv("WORKERS_AMOUNT")

	app.Init()

	mux := api.NewMux()
	handler := middleware.Logging(mux)

	for i := 0; i < getEnvInt(wa, 2); i++ {
		go worker.Run(sp, handler)
	}

	select {}
}

func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)

	n, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return n
}
