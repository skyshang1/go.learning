package main

import "fmt"


func main() {
	testMap := make(map[string]string)
	testMap["123"] = "hello"
	
	name, has := testMap["234"]
	fmt.Println("name: ", name)
	fmt.Printf("has: %+v\n", has)

	var m map[string]string
	if m == nil {
		fmt.Println("m is nil")
	}
	h, _ := m["hello"]
	if h == "" {
		fmt.Println("don't panic")
	}
}
