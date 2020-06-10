package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const (
	CtxRecreateTimes = "RECREATE_TIMES"
)

func main() {
	var times int = 3
	context := make(map[string]interface{})
	context[CtxRecreateTimes] = times

	// marshal to string
	data, _ := json.Marshal(context)
	// unmarshal to context
	newContext := make(map[string]interface{})
	err := json.Unmarshal(data, &newContext)
	if err != nil {
		log.Fatal(err)
	}
	// convert interface{} to int
	x := newContext[CtxRecreateTimes].(int)
	fmt.Println(x)
}
