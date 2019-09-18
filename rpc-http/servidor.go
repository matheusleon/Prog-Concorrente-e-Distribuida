package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"encoding/json"
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
type RCVR int

func (a *RCVR) Msg_Function(msg []byte, ans *[]byte) error {
	msgRequest := Request{}
	err := json.Unmarshal(msg, &msgRequest)
	failOnError(err, "Unmarshal failed")
	r := msgRequest.P1 + msgRequest.P2
	replyMsg := Reply{Result: r}
	replyMsgBytes, err := json.Marshal(replyMsg)
	failOnError(err, "Marshal failed")
	*ans = replyMsgBytes
	return nil
}

func main() {
	RCVR := new(RCVR)
	err := rpc.Register(RCVR)
	failOnError(err, "error registering RCVR")

	rpc.HandleHTTP()

	listener, _ := net.Listen("tcp", ":8080")

	http.Serve(listener, nil)
	failOnError(err, "error serving")
}
