package piscine

func TrimAtoi(s string) int {

	var x int
	letter := 'a'
	sign := 1
	isNumberMet := false
	for index := range s {

		letter = rune(s[index])
		if letter == '-' && !isNumberMet {
			sign = sign * (-1)
			continue
		} else if letter < '0' || letter > '9' {
			continue
		} else {
			isNumberMet = true
			numb := 0
			for i := '0'; i < letter; i++ {
				numb++
			}
			x = x*10 + (numb)
		}

	}
	return x * sign
}
