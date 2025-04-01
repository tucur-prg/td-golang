package main

import (
	"fmt"
)

// [Go]「...（ピリオド３つ）」の使い方（基本編） #初心者 - Qiita
// https://qiita.com/Yashy/items/a02b74a00136dc5a42c4
func main() {
	fmt.Println(sum(1, 2, 3, 4))

	s := []int{1, 2, 3, 4}
	fmt.Println(sum(s...))

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	s3 := append(s1, s2...)

	fmt.Println(s3)
}

func sum(nums ...int) int {
	var res int
	for _, v := range nums {
		res += v
	}
	return res
}