package main

import (
	"fmt"
	"sync"
	"time"
)

// Test Goal:
// validate goroutine

func main() {
	var curMasterIP string
	var mux sync.Mutex
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func(waitGroup *sync.WaitGroup) {
		defer waitGroup.Done()

		mux.Lock()
		fmt.Println(time.Now())
		curMasterIP = "127.0.0.1"
		fmt.Println("curMasterIP in goroutine1:", curMasterIP)
		time.Sleep(10 * time.Second)
		mux.Unlock()
		fmt.Println(time.Now())

	}(&waitGroup)

	go func(waitGroup *sync.WaitGroup) {
		defer waitGroup.Done()

		mux.Lock()
		fmt.Println(time.Now())
		curMasterIP = "127.0.0.100"
		fmt.Println("curMasterIP in goroutine2:", curMasterIP)
		time.Sleep(5 * time.Second)
		mux.Unlock()

	}(&waitGroup)

	waitGroup.Wait()

	fmt.Println(time.Now())
	fmt.Println("curMasterIP in main:", curMasterIP)
}
