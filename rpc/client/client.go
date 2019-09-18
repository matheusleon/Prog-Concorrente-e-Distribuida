package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Request struct {
	P1 int
	P2 int
}

type Reply struct {
	Result int
}

func main() {

	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}
	
	var ans Reply
	for i := 0; i < 10000; i++ {
		start := time.Now()
		
		msg := Request{5, 7}
		client.Call("RCVR.Msg_Function", msg, &ans)
		
		elapsed := time.Since(start)
		
		//fmt.Println(ans)
		fmt.Printf("%s\n", elapsed)	
	}
}




