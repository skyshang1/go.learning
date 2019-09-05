package main

import (
	"net/http"
	"fmt"
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hi, This is an example of https service in golang!")
}

func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil)
}