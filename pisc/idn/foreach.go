package piscine

func ForEach(f func(int), a []int) {
	for _, z := range a {
		f(z)
	}
}
