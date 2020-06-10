package main

// observe gc process:
// go build go_trace.go
// env GODEBUG=gotrace=1 ./go_trace
func allocate() {
	_ = make([]byte, 1<<20)
}

func main() {
	for n := 1; n < 10000; n++ {
		allocate()
	}
}
