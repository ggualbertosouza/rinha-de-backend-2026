package main

import (
	"log"

	loadBalancer "github.com/ggualbertosouza/rinha-de-backend-2026/internal/lb"
)

func main() {
	lb := loadBalancer.NewLoadBalancer("9999", "/tmp/lb.sock")

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
