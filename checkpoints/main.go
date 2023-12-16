package main

import (
	// "fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	result := 0
	n1 := os.Args[1]
	sign := os.Args[2]
	n2 := os.Args[3]
	res := ""
	c1 := 0
	if sign == "+" || sign == "-" || sign == "*" || sign == "/" || sign == "%" {
		c1++
	} else {
		return
	}

	if n2 == "0" && sign == "/" {
		res = "No division by 0"
		os.Stdout.WriteString(res + "\n")
		return
	}
	if n2 == "0" && sign == "%" {
		res = "No modulo by 0"
		os.Stdout.WriteString(res + "\n")
		return
	}
	num1, bool2 := Atoi(n1)
	num2, bool1 := Atoi(n2)
	if bool2 == false || bool1 == false {
		return
	}
	// fmt.Println(num1)
	// fmt.Println(num2)
	if sign == "+" {
		result = num1 + num2
	} else if sign == "-" {
		result = num1 - num2
	} else if sign == "*" {
		result = num1 * num2
	} else if sign == "/" {
		result = num1 / num2
	} else if sign == "%" {
		result = num1 % num2
	}
	res = itoa(result)
	os.Stdout.WriteString(res + "\n")
}

func Atoi(s string) (int, bool) {
	result := 0
	a := 0
	digit := 1
	n := 0
	for _, val := range s {
		if val >= '0' && val <= '9' || val == '+' || val == '-' {
			if val == '+' && a == 0 {
				a = 1
				// fmt.Println(a)
				continue
			} else if val == '-' && a == 0 {
				a = 1
				digit = -1
				continue
			}
			if val >= '0' && val <= '9' {
				a = 1
			}
			if val == '-' || val == '+' {
				if a != 0 {
					return 0, false
				}
			}
			n = result
			result = result*10 + int(val-'0')
			if n != 0 && result/n < 0 {
				return 0, false
			}
		} else {
			return 0, false
		}
	}
	if result*digit >= 9223372036854775807 || result*digit <= -9223372036854775808 {
		return 0, false
	}
	return result * digit, true
}

func itoa(n int) string {
	s := ""
	sign := false
	if n == 0 {
		return "0"
	}
	if n < 0 {
		sign = true
		n = -n
	}
	for n > 0 {
		s = string(n%10+'0') + s
		n /= 10
	}
	if sign {
		return "-" + s
	}
	return s
}
