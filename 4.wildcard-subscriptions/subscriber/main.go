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

	sub1, _ := nc.Subscribe("orders.created", func(msg *nats.Msg) {
		log.Printf("subject: %s\n", msg.Subject)
		log.Printf("message received: %s\n", string(msg.Data))
		log.Println("=======================================")
	})
	defer sub1.Unsubscribe()

	sub2, _ := nc.Subscribe("orders.updated", func(msg *nats.Msg) {
		log.Printf("subject: %s\n", msg.Subject)
		log.Printf("message received: %s\n", string(msg.Data))
		log.Println("=======================================")
	})
	defer sub2.Unsubscribe()

	sub3, _ := nc.Subscribe("orders.*", func(msg *nats.Msg) {
		log.Printf("wildcard: orders.*")
		log.Printf("subject: %s\n", msg.Subject)
		log.Printf("message received: %s\n", string(msg.Data))
		log.Println("=======================================")
	})
	defer sub3.Unsubscribe()

	sub4, _ := nc.Subscribe("orders.*.created", func(msg *nats.Msg) {
		log.Printf("wildcard: orders.*.created")
		log.Printf("subject: %s\n", msg.Subject)
		log.Printf("message received: %s\n", string(msg.Data))
		log.Println("=======================================")
	})
	defer sub4.Unsubscribe()

	sub5, _ := nc.Subscribe("orders.123.>", func(msg *nats.Msg) {
		log.Printf("wildcard: orders.123.>")
		log.Printf("subject: %s\n", msg.Subject)
		log.Printf("message received: %s\n", string(msg.Data))
		log.Println("=======================================")
	})
	defer sub5.Unsubscribe()

	log.Print("Listening on subjects:\n")

	time.Sleep(60 * time.Second)
}
