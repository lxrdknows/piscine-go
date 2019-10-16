package main

import "github.com/01-edu/z01"

func main() {
	var a = 'a'
	for a <= 'z' {
		z01.PrintRune(rune(a))
		a++
	}

	z01.PrintRune('\n')
}
