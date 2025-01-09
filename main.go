package main

import (
	"fmt"
	"os"
	"symlang/sym"
)

func main() {
	args := os.Args

	if len(args) == 2 {
		runtime := sym.NewRuntime()
		runtime.ExecFile(args[1])
	} else {
		fmt.Println("Missing arguments. Usage: go run main.go <file>.")
		os.Exit(64)
	}
}
