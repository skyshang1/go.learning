package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"crypto/tls"
)

// Ignore certificate validating error
func main() {
	tr := http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: &tr}

	resp, err := client.Get("https://localhost:8081")
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
