package main

import "github.com/01-edu/z01"

func main() {
	var a = '0'
	for a <= '9' {
		z01.PrintRune(rune(a))
		a++
	}

	z01.PrintRune('\n')
}