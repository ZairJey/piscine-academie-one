package piscine

func ActiveBits(n int) int {
	var suka int
	for n > 0 {
		if n%2 == 1 {
			suka++
		}
		n = n / 2
	}
	return suka
}
