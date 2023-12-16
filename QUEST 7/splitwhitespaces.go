package piscine

func SplitWhiteSpaces(s string) []string {
	var ww []string
	var ee string
	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' || char == '\r' {
			if ee != "" {
				ww = append(ww, ee)
				ee = ""
			}
		} else {
			ee += string(char)
		}
	}
	if ee != "" {
		ww = append(ww, ee)
	}
	return ww
}
