package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type Request struct {
	UserId string "json:'user_id'"
}

type Respone struct {
	UserId string "json:'user_id'"
	Name   string "json:'name'"
	Email  string "json:'email'"
}

// replier (server)
func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	sub, err := nc.Subscribe("user.get", func(msg *nats.Msg) {
		var req Request

		json.Unmarshal(msg.Data, &req)

		// simulate DB l0okup
		user := Respone{
			UserId: req.UserId,
			Name:   "NAME",
			Email:  "example@mail.com",
		}

		res, _ := json.Marshal(user)
		msg.Respond(res)
	})

	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	time.Sleep(time.Second * 60)
}
