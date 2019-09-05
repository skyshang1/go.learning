package main

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	Key   string `json:"key,required"`
	Value string `json:"value,required"`
}

func main() {
	var item Item
	itemStr := "{\"key\": \"test\"}"
	err := json.Unmarshal([]byte(itemStr), &item)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(item)
}
