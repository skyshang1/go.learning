package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Format("20060102150405"))
	fmt.Println( time.Now().In(time.Local).Format(time.RFC3339))
}
