package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	args := os.Args[1]
	var result string

	for i := len(args) - 1; i >= 0; i-- {
		if len(result) > 0 && args[i] == ' ' {
			break
		}
		if args[i] != ' ' {
			result = string(args[i]) + result
		}
	}
	if len(result) > 0 {
		for _, qwe := range result {
			z01.PrintRune(qwe)
		}
		z01.PrintRune('\n')
	}
}
