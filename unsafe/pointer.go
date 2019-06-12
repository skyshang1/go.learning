package main

import (
	"unsafe"
	"fmt"
)

type Dog struct {
	name string
	age int
	color string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

// unsafe.Pointer
func main() {
	dog := Dog{name: "little pig"}
	dogP := &dog
	dogPtr := uintptr(unsafe.Pointer(dogP))

	// *dog -> unsafe.Pointer -> uintptr
	// uintptr -> unsafe.Pointer -> *dog
	// *dog <-> unsafe.Pointer
	// unsafe.Pointer <-> uintptr
	namePtr := dogPtr + unsafe.Offsetof(dogP.name)
	nameP := (*string) (unsafe.Pointer(namePtr))
	agePtr := dogPtr + unsafe.Offsetof(dogP.age)
	fmt.Printf("dog: %v, dog.name: %v, dog.age: %v\n", dogPtr, namePtr, agePtr)

	fmt.Println("name: ", *nameP)
}
