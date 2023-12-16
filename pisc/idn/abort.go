package piscine

func Abort(a, b, c, d, e int) int {
	array := []int{a, b, c, d, e}
	array = BubbleSort(array)
	return array[2]
}

func BubbleSort(array []int) []int {
	for x := 0; x < len(array)-1; x++ {
		for y := 0; y < len(array)-x-1; y++ {
			if array[y] > array[y+1] {
				array[y], array[y+1] = array[y+1], array[y]
			}
		}
	}
	return array
}
