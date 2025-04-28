package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

type Sample struct {
	ID   string
	Name string
}

func main() {
	vals := []Sample{
		{ID: "1", Name: "value1"},
		{ID: "2", Name: "value2"},
		{ID: "3", Name: "value3"},
		{ID: "4", Name: "value4"},
	}

	fmt.Println(vals)

	fmt.Println(slices.IndexFunc(vals, func(v Sample) bool { return v.Name == "value3" }))
	fmt.Println(slices.IndexFunc(vals, func(v Sample) bool { return v.Name == "value5" }))
}