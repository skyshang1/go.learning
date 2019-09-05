package main

import "fmt"


func main(){
	ch := make(chan int)
	go func(){
		defer func(){
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
		close(ch)

		// put message to closed channel
		ch <- 10
	}()
	msg, ok := <-ch
	if ok {
		fmt.Println("receive message:", msg)
	} else {
		fmt.Println("channel closed ")
	}
}
