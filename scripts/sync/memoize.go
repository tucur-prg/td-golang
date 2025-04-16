package main

import (
	"fmt"
	"sync"
	"context"
)

func main() {
	ctx := context.Background()

	f := Memoize(call)

	{
		res, err := f(ctx, "100-2002-a011")
		if err != nil {
			panic(err)
		}
		
		fmt.Printf("%T: %+v\n", res, res)
	}
	{
		res, err := f(ctx, "100-2002-a011")
		if err != nil {
			panic(err)
		}
		
		fmt.Printf("%T: %+v\n", res, res)
	}
	{
		res, err := f(ctx, "100-2002-a012")
		if err != nil {
			panic(err)
		}
		
		fmt.Printf("%T: %+v\n", res, res)
	}
}

type Credential struct {
	Id string
	Secret string
}

func call(ctx context.Context, id string) (*Credential, error) {
	fmt.Println(id)
	return &Credential{"00001", "abcdef"}, nil
}

func Memoize[T comparable, R any](f func(context.Context, T) (*R, error)) func(context.Context, T) (*R, error) {
	cache := make(map[T]*R)

	var mu sync.Mutex

	return func(ctx context.Context, arg T) (*R, error) {
		mu.Lock()
		defer mu.Unlock()

		if val, ok := cache[arg]; ok {
			return val, nil
		}

		result, err := f(ctx, arg)
		if err != nil {
			return nil, err
		}

		cache[arg] = result
		return result, nil
	}
}