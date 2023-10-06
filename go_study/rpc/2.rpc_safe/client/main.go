package main

import (
	"fmt"
	"log"
	"net/rpc"
	"rpc/2.rpc_safe/pb"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Call(pb.HelloServiceName+".Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
