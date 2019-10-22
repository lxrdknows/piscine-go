package piscine

func NRune(s string, n int) rune {
	r := []rune(s)
	l := 0
	for range r {
		l++
	}
	if l > 0 && n <= l && n > 0 {
		return r[n-1]
	} else {
		return 0
	}
}
