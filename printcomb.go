package piscine

import "fgithub.com/01-edu/z01"

func PrintComb() {
	next := false
	for a = '0'; a <= '9'; a++ {
		for b := a + 1; b <= '9'; b++ {
			for c := b + 1; c <= '9'; c++ {
				if next {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				next = true
				z01.PrintRune(a)
				z01.PrintRune(b)
				z01.PrintRune(c)
			}
		}
	}
	z01.PrintRune(10)
}
