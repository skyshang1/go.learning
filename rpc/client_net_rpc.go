package main

import (
	"goLearning/rpc/service"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatalf("Error in dialing. %s", err)
	}

	// make arguments object
	args := &service.CalculateArgs{
		A: 2,
		B: 3,
	}
	// this will store returned result
	var result service.CalculateResult
	// call remote procedure with args
	err = client.Call("Calculate.Multiply", args, &result)
	if err != nil {
		log.Fatalf("error in Calculate %s", err)
	}
	//
	log.Printf("%d*%d=%d\n", args.A, args.B, result)
}
