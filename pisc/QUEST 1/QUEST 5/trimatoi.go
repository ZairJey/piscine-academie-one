package piscine

func TrimAtoi(s string) int {
	znak := 1
	digit := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] >= '0' && s[i] <= '9' {
			break
		}
		if s[i] == '-' {
			znak = -1
		}
	}
	for _, val := range s {
		if val >= '0' && val <= '9' {
			val = val - 48
			digit = digit*10 + int(val) // собираем все цифры , пока форик не закончится
		}
	}
	return digit * znak
}
