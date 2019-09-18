package main

import (
	"encoding/json"
	"fmt"
	"log"

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

	msgsFromClient, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	fmt.Println("Server is ready...")
	for d := range msgsFromClient {

		// recv request
		msgRequest := Request{}
		err := json.Unmarshal(d.Body, &msgRequest)
		failOnError(err, "Unmarshal failed")

		r := msgRequest.P1 + msgRequest.P2
		// process request
		// prepare response
		replyMsg := Reply{Result: r}
		replyMsgBytes, err := json.Marshal(replyMsg)
		failOnError(err, "Marshal failed")

		// publica resposta
		err = ch.Publish("", msgs.Name, false, false,
			amqp.Publishing{ContentType: "text/plain", Body: []byte(replyMsgBytes)})
		failOnError(err, "Failed to publish a message")
	}
}
