package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Message struct {
	Text string
}

func main() {

	client, err := rpc.DialHTTP("tcp", "localhost:8080")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	
	msg := Message{"Hello Server!"}
	var reply Message
	client.Call("API.Msg_Function", msg, &reply)
	
	fmt.Println(reply)

}
