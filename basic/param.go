package main

import (
	"fmt"
)

func Test1(strs ...string) {
	fmt.Println("Test1: ", strs)
	Test2(strs...)
}

func Test2(strs ...string) {
	fmt.Println("Test2: ", strs)
}

func main() {
	Test1("hello", "world", "test")
}
