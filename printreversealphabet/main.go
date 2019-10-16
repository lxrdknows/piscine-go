package main

import "github.com/01-edu/z01"
func main() {
	var a = 'z'
	for a >= 'a' {
		z01.PrintRune(a)
		a--
	}

	z01.PrintRune(10)
}
