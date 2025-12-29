package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	stream, err := js.AddStream(&nats.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"orders.*"},
		Storage:  nats.FileStorage,
		MaxAge:   24 * time.Hour,
		MaxBytes: 1024 * 1024 * 1024, //1GB
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Created straem: %s", stream.Config.Name)
}
