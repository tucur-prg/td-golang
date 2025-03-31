package main

import (
	"fmt"
)

func main() {
	// interface{}型の値がstring型であるかどうかを確認する場合は、以下のように書きます。
	var x interface{} = "Hello"
	s, ok := x.(string) // アサーション
	if !ok {
		panic("x is not a string")
	}	

	fmt.Printf("%T: %v\n", s, s) // Hello
}
