package piscine

func AlphaCount(str string) int {
	counter := 0

	for _, gg := range str {
		if gg >= 'A' && gg <= 'Z' {
			counter++
		} else if gg >= 'a' && gg <= 'z' {
			counter++
		}
	}
	return counter
}
