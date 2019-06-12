package main

import (
	"fmt"
	"os/signal"
	"time"
)

func main() {
	done := make(chan bool, 1)

	signal.Ignore()
	// This goroutine executes a blocking receive for
	// signals. When it gets one it'll print it out
	// and then notify the program that it can finish.
	go func(){
		time.Sleep(300 * time.Second)
		done <- true
	}()

	// The program will wait here until it gets the
	// expected signal (as indicated by the goroutine
	// above sending a value on `done`) and then exit.
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
