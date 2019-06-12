package main

import (
	"fmt"
	"path"
)


func main() {
	workDir := path.Join("jvessel", "log", "filebeat")
	fmt.Println("Work Dir: ", workDir)
	configPath := path.Join("jvessel", "filebeat.yml")
	fmt.Println("Config Path: ", configPath)
}
