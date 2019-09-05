package main

import (
	"fmt"
	"os"
	"github.com/google/uuid"
	"io/ioutil"
	"strings"
	"bytes"
	"os/exec"
	"path/filepath"
)

func GetDeviceUUID(path string) string {
	// Find MountDir
	mountDir, err := findMountDir(path)
	if err != nil {
		fmt.Printf("Find MountDir error: %s\n", err.Error())
		return ""
	}
	//
	uuidPath := filepath.Join(mountDir, ".uuid")
	_, err = os.Stat(uuidPath)
	if os.IsNotExist(err) {
		// assign uuid
		uuid := uuid.New().String()
		// write to file
		fmt.Println("******UUID FilePath:", uuidPath, "*******")
		err1 := ioutil.WriteFile(uuidPath, []byte(uuid), 0644)
		if err1 != nil {
			fmt.Printf("Write UUID to file error: %s\n", err1.Error())
			return ""
		}
		return uuid
	}

	// read uuid
	if contents,err := ioutil.ReadFile(uuidPath); err == nil {
		uuid := strings.Replace(string(contents),"\n","",1)
		return uuid
	} else {
		fmt.Printf("Read file error: %s\n", err.Error())
	}

	return ""
}


func findMountDir(path string) (string, error) {
	var out bytes.Buffer

	cmd := exec.Command("/bin/bash", "-c",
		fmt.Sprintf("findmnt -n -o TARGET --target %s", path))
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return strings.Replace(out.String(),"\n","",1), err
}


func main() {
	uuid := GetDeviceUUID("/home/skyshang/go/src/github.com/elastic/beats/Makefile")
	if uuid == "" {
		fmt.Println("get device uuid error")
	} else {
		fmt.Println("UUID:", uuid)
	}
}
