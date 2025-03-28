package main

import (
	"fmt"
	"time"
)

func main() {
	taro := &User{ID: "123", Name: "taro"}
	john := &User{ID: "124", Name: "john"}

	hoge := call[*User](taro, john, func(v *User) string { return v.Name })

	fmt.Println("%T: %v\n", hoge, hoge)
}

type User struct {
	ID string
	Name string
	Stamp time.Time
}
func (u User) GetID() string {
	return u.ID
}
func (u *User) SetID(v string) {
	u.ID = v
} 
func (u User) GetStamp() time.Time {
	return u.Stamp
}

type Hoge interface {
	GetID() string
	SetID(string)
	GetStamp() time.Time
}

func call[T Hoge](a, b T, key func(T) string) T {
	fmt.Printf("%T: %v\n", a, a)
	fmt.Printf("%T: %v\n", b, b)
	fmt.Println(key(a), key(b))
	fmt.Println(a.GetID())
	b.SetID("hoge")

	return b
}