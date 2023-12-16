package piscine

func IterativePower(nb int, power int) int {
	if nb > 25 || nb < -25 || power < 0 {
		return 0
	}
	g := 1
	for i := 1; i <= power; i++ {
		g = nb * g
	}
	return g
}
