package main

import (
	"fmt"
)

func main() {
	a := Sample{ID: "123", Name: "abc"}

	a.Update1()
	fmt.Println("Name:", a.Name)

	a.Update2()
	fmt.Println("Name:", a.Name)

	fmt.Println(a.GetDescription1())
	fmt.Println(a.GetDescription2())
}

type Sample struct {
	ID string
	Name string
	Description *string
}

func (s Sample) GetDescription1() string {
	if s.Description == nil {
		return ""
	}
	return *s.Description
}
func (s *Sample) GetDescription2() string {
	if s.Description == nil {
		return ""
	}
	return *s.Description
}

// Pointerじゃないので構造体がコピーされて元の値が変更されない
func (s Sample) Update1() {
	s.Name = "method1"
	fmt.Println("Update1:", s.Name)
}
func (s *Sample) Update2() {
	s.Name = "method2"
	fmt.Println("Update2:", s.Name)
}
