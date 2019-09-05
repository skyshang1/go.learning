package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

// Get https://localhost:8081: x509: certificate signed by unknown authority
func main() {
	resp, err := http.Get("https://localhost:8081")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read Response Body Error:", err)
		return
	}

	fmt.Println(string(body))
}
