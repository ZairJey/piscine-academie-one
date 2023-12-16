package piscine

func ToLower(s string) string {
	var a string
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			char = char + 32
		}
		a = a + string(char)
	}
	return a
}
