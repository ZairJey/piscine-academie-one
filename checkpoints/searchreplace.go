package piscine

import (
	//"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {

	args := os.Args[1:]
	if len(args) != 3 {
		return
	}

	a := os.Args[1]
	b := os.Args[2]
	c := os.Args[3]

	var res string
	for _, q := range a {
		if q == rune(b[0]) {
			q = rune(c[0])
		}
		res += string(q)
	}
	for _, d := range res {
		z01.PrintRune(d)
	}
	z01.PrintRune('\n')
}
