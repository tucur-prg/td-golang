package main

import (
	"errors"
	"fmt"
)

func main() {
	s := "hello"

	{
		fmt.Println(s)
		s := "world"
		fmt.Println(s)
	}
	fmt.Println(s)

	// 変数束縛
	// Go言語の記述の迷いどころについて
	// https://zenn.dev/nobonobo/articles/19c84c231aff46
	if err := call(); err != nil {
		fmt.Println(err)
	}
}

func call() error {
	return errors.New("test")
}