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

func addReplier(idx int, nc *nats.Conn) {
	nc.Subscribe("math.add", func(msg *nats.Msg) {
		var req Request

		json.Unmarshal(msg.Data, &req)

		res := Respone{
			Result: req.FirstOperand + req.SecondOperand,
		}
		resByte, _ := json.Marshal(res)
		log.Printf("Replier %d respond\n", idx)
		msg.Respond(resByte)
	})
}

// replier (server)
func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	for i := 1; i <= 3; i++ {
		addReplier(i, nc)
	}

	time.Sleep(time.Second * 60)
}
