package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpc/3.rpc_json/pb"
)

func main() {
	rpc.RegisterName("HelloService", new(pb.HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
