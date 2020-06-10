package main

import (
	"fmt"
	"runtime"
	"time"
)

// observe stw process
func main() {
	go func() {
		i := 0
		for i < 1000 {
			i++
		}
	}()

	time.Sleep(time.Millisecond)
	runtime.GC()
	fmt.Print("OK")
}
