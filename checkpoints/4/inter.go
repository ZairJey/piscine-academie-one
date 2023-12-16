package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	slice := make([]int, 128)
	if len(arg) != 3 {
		return
	}
	for i := 0; i < len(arg[2]); i++ {
		slice[arg[2][i]] = 1
	}
	fmt.Println(slice)

	for i := 0; i < len(arg[1]); i++ {
		if slice[arg[1][i]] == 1 {
			z01.PrintRune(rune(arg[1][i]))
			slice[arg[1][i]] = 2
		}
	}
	fmt.Println(slice)

	z01.PrintRune('\n')
	fmt.Println(slice)
}
