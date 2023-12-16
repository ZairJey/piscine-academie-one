package piscine

func AlphaCount(s string) int {
	var count int
	for _, elina := range s {
		if elina >= 'a' && elina <= 'z' || elina >= 'A' && elina <= 'Z' {
			count++
		}
	}
	return count
}
