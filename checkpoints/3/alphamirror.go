package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) != 2 {
		//z01.PrintRune('\n')
		return
	}
	digit := ""
	for _, val := range os.Args[1] {
		if val >= 'a' && val <= 'z' {
			digit += string('z' - (val - 'a'))
		} else if val >= 'A' && val <= 'Z' {
			digit += string('Z' - (val - 'A'))
		} else {
			digit += string(val)
		}
	}
	for _, qw := range digit {
		z01.PrintRune(qw)
	}
	z01.PrintRune('\n')
}
