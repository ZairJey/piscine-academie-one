package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	if len(os.Args) != 2 {
		return
	}
	digit := ""
	for _, val := range os.Args[1] {
		if val >= 'a' && val <= 'z' {
			digit += string((val+13-'a')%26 + 'a')
		} else if val >= 'A' && val <= 'Z' {
			digit += string((val+13-'A')%26 + 'A')
		} else {
			digit += string(val)
		}
	}
	for _, ch := range digit {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
}
