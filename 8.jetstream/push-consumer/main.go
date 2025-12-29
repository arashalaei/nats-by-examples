package main

import (
	"encoding/json"
	"log"
	"time"

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

	sub, err := js.Subscribe("order.created", func(msg *nats.Msg) {
		var o order
		json.Unmarshal(msg.Data, &o)
		log.Printf("Received order ID: %s, Amount: %.2f", o.ID, o.Amount)
		msg.Ack()
	}, nats.Durable("ORDER_CONSUMER"), nats.ManualAck())
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	log.Printf("Subscribed to subject: %s", sub.Subject)
	time.Sleep(10 * time.Minute)
}
