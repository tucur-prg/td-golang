package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Printf("%#v\n", ctx)

	ctx = context.WithValue(ctx, "key1", "value1")

	fmt.Printf("%#v\n", ctx)

	ctx = context.WithValue(ctx, "key2", "value2")

	fmt.Printf("%#v\n", ctx)

	fmt.Printf("%T: %v\n", ctx.Value("key1"), ctx.Value("key1"))
}
