package main

import (
	"sync"
	"fmt"
)

func main() {

	var waitGroup sync.WaitGroup
	waitGroup.Add(3)
	// start two goroutines
	var num1, num2, num3 = 100, 200, 300
	go AddNumber(&num1, 10, &waitGroup)
	go AddNumber(&num2, 20, &waitGroup)
	go AddNumber(&num3, 20, &waitGroup)
	// wait all goroutine exit
	waitGroup.Wait()
	// log this info
	fmt.Println("All Goroutine has exited!")
	// verify result
	fmt.Println("num1:", num1, ", num2: ", num2)
}

func AddNumber(number *int, delta int, waitGroup *sync.WaitGroup) {
	fmt.Println("Add Number", *number)
	*number += delta
	waitGroup.Done()
}
