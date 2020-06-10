package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	startTime := time.Now()
	for i := 0; i < 100000; i++ {
		file, err := os.OpenFile("file.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatalln(err)
		}
		file.Close()
	}
	fmt.Println(time.Now().Sub(startTime).Seconds())

	//
	//fmt.Println(file.)

	//fileInfo, err := file.Stat()
	//fmt.Println(fileInfo.Name())
	//fmt.Println(fileInfo.Sys())
}
