package piscine

import "github.com/01-edu/z01"

func PrintNbr(x int) {
	if x < 0 {
		z01.PrintRune('-')
	}
	SetNbr(x)
}

func SetNbr(x int) {
	h := '0'
	if x == 0 {
		z01.PrintRune(h)
		return
	}
	for i := 1; i <= x%10; i++ {
		h++
	}
	for i := -1; i >= x%10; i-- {
		h++
	}
	if x/10 != 0 {
		SetNbr(x / 10)
	}
	z01.PrintRune(h)
	return
}
