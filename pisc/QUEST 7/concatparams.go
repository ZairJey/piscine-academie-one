package piscine

func ConcatParams(args []string) string {
	var zair string
	for i := 0; i < len(args); i++ {
		zair += args[i]
		if i < len(args)-1 {
			zair += "\n"
		}
	}
	return zair
}
