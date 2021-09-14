package main

import (
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	gofakeit.Seed(time.Now().UnixNano())
	go processSignals()
	r := router()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("cannot start the server on :8080 > %s", err)
	}
}
