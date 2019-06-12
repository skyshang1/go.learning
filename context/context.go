package main

import (
	"context"
	"log"
	"time"
	"os"
)

var logger *log.Logger

// WithCancel
func someHandler() {
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)
	time.Sleep(10 * time.Second)
	cancel()
}

func doStuff(ctx context.Context) {
	for {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			logger.Printf("done")
			return
		default:
			logger.Printf("work")
		}
	}
}

// WithTimeout
func timeoutHandler() {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	go doStuff(ctx)
	time.Sleep(10 * time.Second)
	cancel()
}

// WithDeadline
// WithTimeout func(parent Context, timeout time.Duration) (Context, CancelFunc)
// WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
func deadlineHandler() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5 * time.Second))
	go doStuff(ctx)
	time.Sleep(10 * time.Second)
	cancel()
}

//
func main() {

	logger = log.New(os.Stdout, "****", log.Ltime)
	// test Context.WithCancel
	someHandler()
	logger.Printf("down")

	// test Context.WithTimeout
	timeoutHandler()
	logger.Printf("down")

	//
	deadlineHandler()
	logger.Printf("down")


}