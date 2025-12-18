package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	subject := "orders.created"
	data := []byte("Order #12010 created")

	err = nc.Publish(subject, data)
	if err != nil {
		log.Fatal(err)
	}

	nc.Flush() // Ensure message is sents

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Println("message published")
}
