package piscine

func ToUpper(s string) string {
	var res string
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			res += string(char - 32)
		} else {
			res = res + string(char)
		}
	}
	return res
}
