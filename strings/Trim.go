package main

import (
	"fmt"
	"strings"
)


func main() {
	test := " shit god        "
	fmt.Println(strings.Trim(test, " "))
}
