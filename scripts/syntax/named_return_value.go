package main

import (
	"fmt"
)

func main() {
	fmt.Println(call())
}

func call() (x, y int) {
	x = 10
	y = 3

	return
}