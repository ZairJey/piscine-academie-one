package piscine

func Compact(ptr *[]string) int {
	for q := len(*ptr) - 1; q >= 0; q-- {
		if (*ptr)[q] == "" {
			*ptr = append((*ptr)[:q], (*ptr)[q+1:]...)
		}
	}
	return len(*ptr)
}
