package main

import "fmt"

type Test struct {
	A int
	B string
}

func (t *Test) String() string{
	return fmt.Sprintf("%+v", *t)
}

func main() {
	testMap := map[string]*Test{
		"test1": &Test{
			A: 1,
			B: "hello",
		},
		"test2": &Test{
			A: 2,
			B: "world",
		},
	}

	fmt.Printf("%v\n", testMap)
	fmt.Printf("%+v\n", testMap)
}
