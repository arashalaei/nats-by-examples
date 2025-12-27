package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	for i := 0; i < 3; i++ {
		nc.QueueSubscribe("task", "worker", func(msg *nats.Msg) {
			log.Printf("worker %d procesing: %s", i+1, string(msg.Data))
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			log.Printf("worker %d completed: %s", i+1, string(msg.Data))
		})
	}

	time.Sleep(100 * time.Millisecond)
}
