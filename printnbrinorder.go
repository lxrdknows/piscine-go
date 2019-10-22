package piscine

import (
	"github.com/01-edu/z01"
)

func recursion(n int) {
	if n == 0 {
		return
	}
	remainder := n % 10
	runeX := '0'
	for i := 0; i < remainder; i++ {
		runeX++
	}
	recursion(n / 10)
	z01.PrintRune(runeX)

}

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		runes := [20]rune{'-', '9', '2', '2', '3', '3', '7', '2', '0', '3', '6', '8', '5', '4', '7', '7', '5', '8', '0', '8'}
		for _, runeZ := range runes {
			z01.PrintRune(runeZ)
		}
		return
	}
	if n == 0 {
		z01.PrintRune('0')
		return
	} else if n < 0 {
		z01.PrintRune('-')
		n = n * (-1)
	}

	recursion(n)

}
