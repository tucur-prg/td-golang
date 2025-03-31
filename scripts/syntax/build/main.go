package main

import (
	"build/env"
	"fmt"
)

// go run -tags=product main.go
// go run -tags=staging main.go
// go run main.go
func main() {
	fmt.Println(env.GetEnv())
}
