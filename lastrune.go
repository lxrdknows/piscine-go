package piscine

func LastRune(s string) rune {
	count := 0
	for range s {
		count++
	}
	if count <= 0 {
		return '\x00'
	}
	str := []rune(s)
	return str[count-1]

}