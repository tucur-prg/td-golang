package main

import (
	"fmt"
)

func main() {
	dict := make(map[string]string)

	dict["a"] = "123"
	dict["b"] = "456"

	if v, ok := dict["b"]; ok {
		fmt.Println(v)
	}
	if v, ok := dict["c"]; ok {
		fmt.Println(v)
	}

	v := dict["a"]
	fmt.Println(v)
}