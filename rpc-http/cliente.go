
package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
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

func main() {

	client, err := rpc.DialHTTP("tcp", "localhost:8080")
	failOnError(err, "Failed to connect to RPC")
	var ans []byte
	for i := 0; i < 10000; i++ {
		start := time.Now()
		
		msgRequest := Request{i, i}
		msgRequestBytes, err := json.Marshal(msgRequest)
		client.Call("RCVR.Msg_Function", msgRequestBytes, &ans)
		
		response := Reply{}
		err = json.Unmarshal(ans, &response)
		failOnError(err, "Unmarshal failed")
		elapsed := time.Since(start)
		t1 := float64(elapsed.Nanoseconds()) / 1000000
		fmt.Println(t1)
	}
}


