package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.OpenFile("file.log", os.O_RDONLY, 0600)
    if err != nil {
        log.Fatalln(err)
    }

	// 
	//fmt.Println(file.)

	fileInfo, err := file.Stat()
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Sys())
}
