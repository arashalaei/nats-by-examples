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

	subject := "orders.created"

	sub, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		log.Printf("Received: %s\n", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	defer sub.Unsubscribe()

	log.Printf("Listening on subject: %s", subject)

	time.Sleep(60 * time.Second)
}
