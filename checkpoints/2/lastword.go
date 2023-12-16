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
	var ggg string

	for i := len(args) - 1; i >= 0; i-- {
		if len(ggg) > 0 && args[i] == ' ' {
			break
		}
		if args[i] == ' ' {
			break
		}
		if args[i] != ' ' {
			ggg = string(args[i]) + ggg
		}
	}
	if len(ggg) > 0 {
		for _, qwe := range ggg {
			z01.PrintRune(qwe)
		}
		z01.PrintRune('\n')
	}
}
