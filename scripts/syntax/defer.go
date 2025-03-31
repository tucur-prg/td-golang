package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	defer close()

	fmt.Println("end")
}

func close() {
	fmt.Println("close")
}
