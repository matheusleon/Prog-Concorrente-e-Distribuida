package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type Request struct {
	P1 int
	P2 int
}

type Reply struct {
	Result int
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("request", false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.QueueDeclare("response", false, false, false, false, nil)
	failOnError(err, "Failed to declare a reply queue")

	msgsFromServer, err := ch.Consume(msgs.Name, "", true, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	for i := 0; i < 100; i++ {
		start := time.Now()

		msgRequest := Request{P1: i, P2: i}
		msgRequestBytes, err := json.Marshal(msgRequest)
		failOnError(err, "Marshal failed")

		err = ch.Publish("", q.Name, false, false,
			amqp.Publishing{ContentType: "text/plain", Body: msgRequestBytes})
		failOnError(err, "Failed to publish a message")

		x := <-msgsFromServer
		response := Reply{}
		err = json.Unmarshal(x.Body, &response)
		//fmt.Println(response.Result)
		elapsed := time.Since(start)
		t1 := float64(elapsed.Nanoseconds()) / 1000000
		fmt.Println(t1)
	}
}
