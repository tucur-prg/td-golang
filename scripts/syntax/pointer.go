package main

import (
	"fmt"
)

func main() {
	a := new(int)
	fmt.Printf("%T %p: %v\n", a, a, *a)

	p := NewPerson()
	fmt.Printf("%T %p: %+v\n", p, p, p)
}

type Person struct {
	Name string
	Age int
}

func NewPerson() *Person {
	return &Person{"Taro", 10}
}
