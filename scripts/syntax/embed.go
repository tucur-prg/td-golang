package main

import (
	"embed" // 1つのファイルを読み出す場合は `_ "embed"`
	"fmt"
)

//go:embed assets/sample.json
var data []byte

//go:embed assets/*.json
var fs embed.FS

func main() {
	fmt.Println(string(data))

	example, err := fs.ReadFile("assets/example.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(example))
}
