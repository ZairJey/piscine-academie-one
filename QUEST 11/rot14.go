package piscine

func Rot14(s string) string {
	var netto string
	for _, val := range s {
		if val >= 'a' && val <= 'z' {
			netto += string((val+14-'a')%26 + 'a')
		} else if val >= 'A' && val <= 'Z' {
			netto += string((val+14-'A')%26 + 'A')
		} else {
			netto += string(val)
		}
	}
	return netto
}
