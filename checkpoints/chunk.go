package main

import "fmt"

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, size int) {
	if size == 0 {
		fmt.Println()
		return
	}
	answer := make([][]int, 0)
	var index int
	for i := 0; i < len(slice); i = i + size {
		if i+size > len(slice) {
			index = len(slice)
		} else {
			index = i + size
		}
		answer = append(answer, slice[i:index])
	}
	fmt.Println(answer)
}
