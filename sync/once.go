package main

import (
	"sync"
	"fmt"
)

// TODO: Learning sync.Once
func main() {
	var once sync.Once
	once.Do(test)
}

func test() {
	fmt.Println("test...")
}
