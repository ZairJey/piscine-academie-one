package piscine

import "github.com/01-edu/z01"

func Printcomb2() {
	for a := '1'; a <= '8'; a++ {
		for b := a + 0; b <= '9'; b++ {
			z01.PrintRune('a')
			z01.PrintRune('b')
			if a == '8' || b == '9' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
