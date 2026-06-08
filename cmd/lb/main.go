package main

import (
	"log"
	"os"

	loadBalancer "github.com/ggualbertosouza/rinha-de-backend-2026/internal/lb"
)

func main() {
	serverPort := os.Getenv("LB_PORT")
	socketPath := os.Getenv("LB_SOCKET")

	lb := loadBalancer.NewLoadBalancer(serverPort, socketPath)

	go func() {
		if err := lb.ListenTcp(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := lb.ListenUnix(); err != nil {
			log.Fatal(err)
		}
	}()

	select {}
}
