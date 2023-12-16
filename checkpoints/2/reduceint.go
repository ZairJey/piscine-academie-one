package main

import (
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}

func ReduceInt(a []int, f func(int, int) int) {

	res := 0
	for i := 0; i < len(a)-1; i++ {
		if i == 0 {
			res = a[0]
		}
		res = f(res, a[i+1])
	}

	q := strconv.Itoa(res)
	for _, i := range q {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}

// func Itoa(val int) string {
// 	var s string
// 	flag := false
// 	if val < 0 {
// 		flag = true
// 		val = -val
// 	}
// 	if val == 0 {
// 		return "0"
// 	}

// 	for val != 0 {
// 		digit := val % 10
// 		s = string(digit+'0') + s
// 		val = val / 10
// 	}
// 	if flag {
// 		return "-" + s
// 	}
// 	return s
// }
