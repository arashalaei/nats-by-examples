package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	var tasks [][]byte

	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("Task number #%d", i+1)
		tasks = append(tasks, []byte(msg))
	}

	for _, t := range tasks {
		err := nc.Publish("task", t)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Print("tasks sent.")
}
