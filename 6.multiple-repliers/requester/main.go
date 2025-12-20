package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type Request struct {
	FirstOperand  int64  `json:"first_operand"`
	Operation     string `json:"operation"`
	SecondOperand int64  `json:"second_Operand"`
}

type Respone struct {
	Result int64 `json:"result"`
}

// requester (client)
func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	reqs := []Request{
		{FirstOperand: 1, Operation: "+", SecondOperand: 2},
		{FirstOperand: 2, Operation: "+", SecondOperand: 3},
		{FirstOperand: 3, Operation: "+", SecondOperand: 4},
		{FirstOperand: 4, Operation: "+", SecondOperand: 5},
		{FirstOperand: 5, Operation: "+", SecondOperand: 6},
		{FirstOperand: 6, Operation: "+", SecondOperand: 7},
		{FirstOperand: 7, Operation: "+", SecondOperand: 8},
		{FirstOperand: 8, Operation: "+", SecondOperand: 9},
		{FirstOperand: 9, Operation: "+", SecondOperand: 10},
		{FirstOperand: 10, Operation: "+", SecondOperand: 1},
	}
	for _, req := range reqs {
		data, _ := json.Marshal(req)
		resByte, err := nc.Request("math.add", data, 2*time.Second)
		if err != nil {
			log.Fatal(err)
		}
		var res Respone
		json.Unmarshal(resByte.Data, &res)
		log.Printf("res: %+v", res)
	}
}
