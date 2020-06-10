package main

import (
	"fmt"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
	fmt.Println("Get Query at", time.Now())
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8888", nil)
}
