package main

import (
	"flag"
	"fmg"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("command not found.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		cmd := flag.NewFlagSet("create", flag.ExitOnError)
		var (
			s = cmd.String("s", "default message.", "string flag")
			i = cmd.Int("i", 0, "int flag")
			b = cmd.Bool("b", false, "bool flag")
		)
		cmd.Parse(os.Args[2:])

		fmt.Println(*s, *i, *b)
	default:
		fmt.Println("command not found.")
		os.Exit(1)
	}
}