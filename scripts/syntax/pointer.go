package main

import (
	"fmt"
)

func main() {
	a := new(int)
	fmt.Printf("%T %p: %v\n", a, a, *a)

	p := NewPerson()
	fmt.Printf("%T %p: %+v\n", p, p, p)

	q := p
	fmt.Printf("%T %p: %+v\n", q, q, q)

	q.Name = "John"
	p.Age = 20
	fmt.Printf("%+v == %+v\n", p, q)
}

type Person struct {
	Name string
	Age int
}

func NewPerson() *Person {
	return &Person{"Taro", 10}
}
