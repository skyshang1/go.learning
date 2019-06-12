package main

import (
	"bytes"
	"github.com/gorilla/rpc/json"
	"goLearning/rpc/service"
	"log"
	"net/http"
)

func main() {
	url := "http://localhost:1234/rpc"
	args := &service.HelloArgs{
		Who: "cuixiaoyu",
	}

	message, err := json.EncodeClientRequest("HelloService.Say", args)
	if err != nil {
		log.Fatalf("%s", err)
	}

	log.Printf("Request: %s\n", string(message))
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(message))
	if err != nil {
		log.Fatalf("%s", err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error in sending request to %s. %s", url, err)
	}
	defer resp.Body.Close()

	var result service.HelloReply

	err = json.DecodeClientResponse(resp.Body, &result)
	if err != nil {
		log.Fatalf("Couldn't decode response. %s", err)
	}
	log.Printf("Reply is: %s\n", result)
}
