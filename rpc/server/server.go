package main

import (
	"log"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Request struct {
	P1 int
	P2 int
}

type Reply struct {
	Result int
}
type RCVR int

func (a *RCVR) Msg_Function(msg Request, ans *Reply) error {
	fmt.Println(msg)
	*ans = Reply{msg.P1 + msg.P2}
	return nil
}

func main() {
	RCVR := new(RCVR)
	err := rpc.Register(RCVR)
	if err != nil {
		log.Fatal("error registering RCVR", err)
	}

	rpc.HandleHTTP()

	listener, _ := net.Listen("tcp", ":8080")

	http.Serve(listener, nil)
	if err != nil {
		log.Fatal("error serving: ", err)
	}
}



