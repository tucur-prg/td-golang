package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	defer close("main")

	call("1")
	call("2")
	call("3")

	fmt.Println("end")
}

func close(id string) {
	fmt.Println("close", id)
}

func call(id string) {
	msg := "call " + id
	defer close(msg)

	fmt.Println(msg)
}
