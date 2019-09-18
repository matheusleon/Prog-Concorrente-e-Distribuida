package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Message struct {
	Text string
}

func main() {

	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	
	var reply Message
	for i := 0; i < 10000; i++ {
		start := time.Now()
		
		msg := Message{"Hello Server!"}
		client.Call("API.Msg_Function", msg, &reply)
		
		elapsed := time.Since(start)
		
		//fmt.Println(reply)
		fmt.Printf("%s\n", elapsed)	
	}
}
