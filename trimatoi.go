package piscine

func TrimAtoi(s string) int {
	bool1 := false // not negative
	arrayStr := []rune(s)
	n := 0 // lenght
	for range arrayStr {
		n++
	}
	if n != 0 && arrayStr[0] == '-' { //if real negative, set negative
		bool1 = true
	}
	ans := 0
	for i := 0; i < n; i++ {
		if arrayStr[i] >= '0' && arrayStr[i] <= '9' {
			ans *= 10
			ans += intfor(arrayStr[i])
		} else if arrayStr[i] == '-' && ans == 0 {
			bool1 = true
		}
	}
	if bool1 == true {
		ans = ans * -1
	}
	return ans
}
