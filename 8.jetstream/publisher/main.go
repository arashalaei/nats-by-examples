package main

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

type order struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amont"`
}

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

	order := order{
		ID:     "ODR#123",
		Amount: 233,
	}

	orderBytes, _ := json.Marshal(order)

	pubAck, err := js.Publish("order.created", orderBytes)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Published message, sequence: %d, stream: %s", pubAck.Sequence, pubAck.Stream)
}
