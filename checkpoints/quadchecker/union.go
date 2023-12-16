package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	p := make([]int, 128)

	if len(args) != 3 {
		return
	}
	args[1] += args[2]
	for i := 0; i < len(args[1]); i++ {
		p[args[1][i]] = 1
	}
	for i := 0; i < len(args[1]); i++ {
		if p[args[1][i]] == 1 {
			z01.PrintRune(rune(args[1][i]))
			p[args[1][i]] = 2
		}
	}
	z01.PrintRune('\n')
}
