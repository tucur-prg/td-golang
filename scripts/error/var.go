package main

import (
	"errors"
	"fmt"
)

func main() {
	err := call(2)
	if err != nil {
		fmt.Println("ErrNum1:", errors.Is(err, ErrNum1))
		fmt.Println("ErrNum2:", errors.Is(err, ErrNum2))
		fmt.Println("ErrNum3:", errors.Is(err, ErrNum3))
	}
}

// Goの標準パッケージでも採用されています。(sql.ErrNoRowsとか)
var (
	ErrNum1 = errors.New("err num 1")
	ErrNum2 = errors.New("err num 2")
	ErrNum3 = errors.New("err num 3")
)

func call(num int) error {
	if num == 1 {
		return ErrNum1
	} else if num == 2 {
		return ErrNum2
	} else if num == 3 {
		return ErrNum3
	}

	return nil
}
