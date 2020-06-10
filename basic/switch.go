package main

import "fmt"

func main() {
	index := 100

	switch index {
	case 100:
		fmt.Println("got 100")
	case 101:
		fmt.Println("got 101")
	default:
		fmt.Println("got default")
	}
}
