package main

import (
	"os/exec"
	"bytes"
	"log"
	"time"
	"fmt"
	"github.com/google/uuid"
)

func exec_shell(s string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), err
}

func main() {
	startTime := time.Now()
	for i := 0; i < 1000; i++ {
		_, err := exec_shell("findmnt -n -o SOURCE --target /home/skyshang/go/src/go.learning/README.md")
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(dev)
	}
	endTime := time.Now()
	fmt.Println(endTime.Sub(startTime).Seconds())

	fmt.Println(uuid.New().String())
}
