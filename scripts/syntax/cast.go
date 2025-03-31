package main

import (
	"strconv"
	"fmt"
)

func main() {
	{
		var i int = 10
		var f float64 = float64(i) // 型変換
		fmt.Printf("%T: %v\n", f, f) // 10.0
	}
	{
		var s string = "42"
		var i int
		i, err := strconv.Atoi(s) // 文字列をintに変換
		if err != nil {
			panic(err)
		}

		fmt.Printf("%T: %v\n", i, i) // 42
	}
}
