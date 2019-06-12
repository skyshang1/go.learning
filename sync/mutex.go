package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Node struct {
	data int
	next *Node
}

type NodeWithLock struct {
	data int
	next *NodeWithLock
	lock sync.Mutex
}

var (
	head = NodeWithLock{data: 0}
	x    = float64(1.0)
)

func main() {
	// init list
	curPosition := &head
	for i := 0; i < 10000; i++ {
		node := NodeWithLock{data: i}
		curPosition.next = &node
		curPosition = curPosition.next
	}

	fmt.Printf("| Element Count | Elapsed Time NoLock | Elapsed Time WithLock |\n")
	for times := 10; times != 1000000; times *= 10 {
		BenchmarkVisitList(times)
	}
}

func BenchmarkVisitList(times int) {
	now := time.Now()
	for index := 0; index < times; index++ {
		VisitList(head)
	}
	elapsedTime := time.Now().Sub(now).Seconds()

	now = time.Now()
	for index := 0; index < times; index++ {
		VisitListWithLock(head)
	}
	elapsedTimeWithLock := time.Now().Sub(now).Seconds()

	fmt.Printf("|\t%vw\t|\t%vs\t|\t%v\t|\n", times, elapsedTime, elapsedTimeWithLock)
}

func VisitList(head NodeWithLock) {
	for p := &head; p != nil; p = p.next {
		x = rand.Float64()
	}
}

func VisitListWithLock(head NodeWithLock) {
	p := &head
	for {
		if p.next == nil {
			return
		}
		p.next.lock.Lock()
		p = p.next
		x = rand.Float64()
		p.lock.Unlock()
	}
}

//| Element Count | Elapsed Time NoLock | Elapsed Time WithLock |
//|	10w	|	0.000239797s	|	0.008102965	|
//|	100w	|	0.003656796s	|	0.021185253	|
//|	1000w	|	0.029801212s	|	0.203868091	|
//|	10000w	|	0.226212125s	|	2.009825262	|
//|	100000w	|	2.219181445s	|	20.217361036	|

// visit with random float64
//| Element Count | Elapsed Time NoLock | Elapsed Time WithLock |
//|	10w	|	0.004640235s	|	0.005169825	|
//|	100w	|	0.035533277s	|	0.062905586	|
//|	1000w	|	0.339447052s	|	0.487624831	|
//|	10000w	|	3.122300947s	|	4.983750482	|
//|	100000w	|	31.558339095s	|	49.357932278	|
