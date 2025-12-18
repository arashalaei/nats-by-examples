package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type Order struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	subject := "orders.created"

	sub, err := nc.Subscribe(subject, func(msg *nats.Msg) {
		var order Order
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Printf("Error: %v", err)
			return
		}
		log.Printf("Received order: %+v", order)
	})
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	time.Sleep(60 * time.Second)
}
