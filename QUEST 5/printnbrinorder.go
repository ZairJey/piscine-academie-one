package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	a := 0
	slice := []int{}
	for n > 0 {
		a = n % 10
		n /= 10
		slice = append(slice, a)
	}
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	for _, val := range slice {
		val = val + '0'
		z01.PrintRune(rune(val))
	}
}
