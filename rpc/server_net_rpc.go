package main

import (
	"goLearning/rpc/service"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	calculator := new(service.Calculate)
	err := rpc.Register(calculator)
	if err != nil {
		log.Fatalf("Format of service Calculate isn't correct. %s", err)
	}
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatalf("Couldn't start listening on port 1234. Error %s", e)
	}
	log.Println("Serving RPC handler")
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatalf("Error serving: %s", err)
	}
}
