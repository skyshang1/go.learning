package main

import (
	"fmt"
	"reflect"
)

const TagName = "Test"

type Info struct {
	Name string `Test:"-"`
	Age  int    `Test:"age,min=17,max=60"`
	Sex  string `Test:"sex,required"`
}

func main() {
	info := Info{
		Name: "ben",
		Age:  23,
		Sex:  "male",
	}

	t := reflect.TypeOf(info)
	fmt.Println("Type:", t.Name())
	fmt.Println("Kind:", t.Kind())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(TagName)
		fmt.Printf("%d. %v(%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
	}
}
