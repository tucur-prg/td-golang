package main

import (
	"fmt"
)

func main() {
	p := NewPerson(182.0, 75.5)
	fmt.Println(p) // &{182 75.5}
}

type person struct {
	height float32
	weight float32
}

// コンストラクタのようなものは実装できます。
// その時の関数名として通例、 New + 構造体名 が適用されます。
func NewPerson(height, weight float32) *person{
	return &person{height: height, weight: weight}
}