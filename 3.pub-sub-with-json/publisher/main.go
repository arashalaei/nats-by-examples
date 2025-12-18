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
	order := Order{
		Id:        "ORD-1234",
		UserId:    "UST-5678",
		Amount:    99.99,
		Status:    "created",
		CreatedAt: time.Now(),
	}

	data, _ := json.Marshal(order)

	err = nc.Publish(subject, data)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("order published")
}
