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

	nc.Publish("orders.created", []byte("Order created."))
	nc.Publish("orders.updated", []byte("Order updated."))

	nc.Publish("orders.123.created", []byte("Order created."))
	nc.Publish("orders.123.updated", []byte("Order created."))
	nc.Publish("orders.456.created", []byte("Order updated."))
	nc.Publish("orders.789.created", []byte("Order created."))

	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Println("messages are published")
}
