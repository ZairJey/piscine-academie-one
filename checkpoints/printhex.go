package main

func main(n int) {

	qwe := "0123456789abcdef"
	ans := ""
	for n > 0 {
		ans = qwe[n%16] + ans
		n = n / 16
	}
	return
}
