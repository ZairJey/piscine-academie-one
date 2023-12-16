package main

import (
	//"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 (
		return
	)
	var result int
	
}







































// 	if len(os.Args) != 4 {
// 		return
// 	}
// 	var result int
// 	var result1 string
// 	val1 := os.Args[1]
// 	sign := os.Args[2]
// 	val2 := os.Args[3]

// 	a := Atoi(val1)
// 	b := Atoi(val2)

// 	if sign == "+" {
// 		result = a + b
// 	} else if sign == "-" {
// 		result = a - b
// 	} else if sign == "*" {
// 		result = a * b
// 	} else if sign == "/" {
// 		result = a / b
// 	} else if sign == "%" {
// 		result = a % b
// 	}
// 	if val2 == "0" && sign == "/" {
// 		result1 = "No division by 0"
// 		os.Stdout.WriteString(result1 + "\n")
// 		return
// 	}
// 	if val2 == "0" && sign == "%" {
// 		result1 = "No modulo by 0"
// 		os.Stdout.WriteString(result1 + "\n")
// 		return
// 	}
	

// 	os.Stdout.WriteString(Itoa(result))
// 	os.Stdout.WriteString("\n")

// }

// func Atoi(s string) int {

// 	digit := 0
// 	count := 1
// 	for _, q := range s {
// 		if q >= '0' && q <= '9' {
// 			digit = digit*10 + int(q-'0')
// 		} else if q == '-' {
// 			count = -1
// 		} else {
// 			return 0
// 		}
// 	}
// 	return (digit * count)
// }

// func Itoa(s int) string {
// 	var digit string
// 	sign := false
// 	if s == 0 {
// 		return "0"
// 	}
// 	if s < 0 {
// 		sign = true
// 		s = -s
// 	}
// 	for s > 0 {
// 		digit = string(s%10+'0') + digit
// 		s /= 10
// 	}
// 	if sign {
// 		return "-" + digit
// 	}
// 	return digit
// }
