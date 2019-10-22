package piscine

import (
	"github.com/01-edu/z01"
)

func recursiveInttoRune(v int, r *[]rune) { //parse int to rune's slice
	c := '0'
	for i := 1; i <= v%10; i++ {
		c++
	}
	if v/10 != 0 {
		recursiveInttoRune(v/10, r)
	}
	*r = append(*r, c)
	return
}

func SwapRune(a *rune, b *rune) {
	c := *b
	*b = *a
	*a = c
}

func sort(r []rune) {
	temp := r
	n := 0
	for range r {
		n++
	}
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if temp[j] < temp[j-1] {
				SwapRune(&temp[j], &temp[j-1])
			}
		}
	}
}

func PrintRuneSliceNnr(n []rune) {
	a := 0
	for index := range n {
		a = index
	}
	for i := 0; i <= a; i++ {
		z01.PrintRune(n[i])
	}

}

func PrintNbrInOrder(n int) {
	r := []rune{}
	recursiveInttoRune(n, &r)
	sort(r)
	PrintRuneSliceNnr(r)
}
