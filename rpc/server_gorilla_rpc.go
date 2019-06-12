package main

import (
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"goLearning/rpc/service"
	"net/http"
)

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(service.HelloService), "")
	http.Handle("/rpc", s)
	http.ListenAndServe(":1234", nil)
}
