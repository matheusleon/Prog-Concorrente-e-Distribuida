package main

import (
	"log"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Message struct {
	Text string
}

type API int

func (a *API) Msg_Function(msg Message, reply *Message) error {
	fmt.Println(msg)
	*reply = Message{"Hello Client!"}
	return nil
}


func main() {
	api := new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("Listener error", err)
	}
	
	log.Printf("serving rpc on port %d", 8080)
	http.Serve(listener, nil)

	if err != nil {
		log.Fatal("error serving: ", err)
	}

}
