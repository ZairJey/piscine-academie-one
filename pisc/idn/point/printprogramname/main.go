package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	po := os.Args[0]
	for _, char := range po[2:] {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
