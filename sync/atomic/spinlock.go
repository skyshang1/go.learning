package atomic

import (
	"sync/atomic"
	"fmt"
	"time"
)

func main() {
	//
}


func spinLock(){
	var lock int32 = 0
	for {
		if atomic.CompareAndSwapInt32(&lock, 0, 1) {
			fmt.Println("")
			break
		}
		time.Sleep(time.Millisecond * 500)
	}
}