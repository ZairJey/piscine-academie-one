package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	wh := os.Args[1:]
	for _, i := range wh {
		for _, a := range i {
			z01.PrintRune(a)
		}
		z01.PrintRune('\n')
	}
}
