package main

import (
	"fmt"
)

func main() {
	n := 2
	switch n {
	case 3:
		fmt.Println("three")
	case 2:
		fmt.Println("two")
	case 1:
		fmt.Println("one")
	}
}
