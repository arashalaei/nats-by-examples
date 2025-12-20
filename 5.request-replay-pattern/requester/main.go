package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type Request struct {
	UserId string `json:"user_id"`
}

type Respone struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

// requester (client)
func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	reqBytes := Request{UserId: "user_123"}
	data, _ := json.Marshal(reqBytes)

	resByte, err := nc.Request("user.get", data, 2*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	var res Respone
	json.Unmarshal(resByte.Data, &res)
	log.Printf("res: %+v", res)
}
