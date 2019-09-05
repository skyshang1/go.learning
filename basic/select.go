package main


import (
	"fmt"
)

func main() {
	resultChan := make(chan int, 10)
	resultChan <- 1
	resultChan <- 2

	for i := 0; i < 2; i++ {
		select {
		case result := <-resultChan:
			fmt.Println(result)
		}
	}
	fmt.Println("select complete...")
}
