package atomic

import (
	"sync/atomic"
	"fmt"
)

// atomic operation
func main() {
	// add
	var x int32 = 100
	fmt.Println("origin value: ", x)
	atomic.AddInt32(&x, 1)
	fmt.Println("value after add: ", x)
	atomic.AddInt32(&x, -1)
	fmt.Println("value after sub: ", x)
	// compare and swap
	isChangeable := atomic.CompareAndSwapInt32(&x, 102, 200)
	fmt.Println("value after cas: ", x, " isChangeable: ", isChangeable)
	// load
	value := atomic.LoadInt32(&x)
	fmt.Println("value loaded:", value)
	// store
	atomic.StoreInt32(&x, 200)
	fmt.Println("value after store:", x)
	// swap
	oldValue := atomic.SwapInt32(&x, 300)
	fmt.Println("new value: ", x)
	fmt.Println("old value: ", oldValue)
}
