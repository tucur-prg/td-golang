package main

import (
	"fmt"
)

func main() {
	// []User
	var list []User
	list = append(list, User{ID: "123", Name: "Hoge"})
	list = append(list, User{ID: "143", Name: "Fuga"})

	fmt.Println(list)

	copy := list

	obj1 := list[0] // 値コピー
	obj2 := &list[0] // ポインタコピー

	fmt.Println(copy)

	for i, _ := range list {
		list[i].Name = "Upate!!"
	}

	fmt.Println(list, copy) // copy も変更されている
	fmt.Println(obj1, obj2)

	// []string
	var strlist []string
	strlist = append(strlist, "hoge")
	strlist = append(strlist, "fuga")

	fmt.Println(strlist)

	strcopy := strlist
	
	fmt.Println(strcopy)

	for i, _ := range strlist {
		strlist[i] = "Upate!!"
	}

	fmt.Println(strlist, strcopy) // strcopy も変更されている
}

type User struct {
	ID string
	Name string
}