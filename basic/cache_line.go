package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Test Result:
// operate normal struct total time elapsed: 22512149ns
// operate optimized struct total time elapsed: 4304262ns

// Conclusion:
// make cpu can use the cache line more efficiently, could improve the operation efficiency of golang.

// References:
// https://segmentfault.com/a/1190000021434918

type Foo struct {
	a uint64
	b uint64
}

// append padding data after valid data, then make cpu can use the cache line more efficiently
type FooWithPadding struct {
	a uint64
	_ [56]byte
	b uint64
	_ [56]byte
}

func main() {
	foo := Foo{}
	timeElapsed := CalculateTimeElapsed(&foo.a, &foo.b)

	fooWithPadding := FooWithPadding{}
	timeElapsedForCacheLine := CalculateTimeElapsed(&fooWithPadding.a, &fooWithPadding.b)

	fmt.Println("operate normal struct total time elapsed:", timeElapsed)
	fmt.Println("operate optimized struct total time elapsed:", timeElapsedForCacheLine)
}

func CalculateTimeElapsed(a *uint64, b *uint64) (timeElapsed int64) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	startTime := time.Now()
	go func(waitGroup *sync.WaitGroup) {
		defer waitGroup.Done()

		for i := 0; i < 1000*1000; i++ {
			atomic.AddUint64(a, 1)
		}
	}(&waitGroup)

	go func(waitGroup *sync.WaitGroup) {
		defer waitGroup.Done()

		for i := 0; i < 1000*1000; i++ {
			atomic.AddUint64(b, 1)
		}
	}(&waitGroup)

	waitGroup.Wait()
	return time.Now().Sub(startTime).Nanoseconds()
}
