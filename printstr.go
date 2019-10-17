package piscine

import "github.com/01-edu/z01"

func PrintStr(str string) {

	strNub := []rune(str)
	for i := range strNub {
		z01.PrintRune(strNub[i])

	}

}
