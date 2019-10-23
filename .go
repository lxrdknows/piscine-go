package main

import "fmt"

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 15 {
		return 0
	}
	a := 1
	b := nb
	for nb := 1; nb <= b; nb++ {

		a = a * nb
	}
	return a
}

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
