package piscine

func StrLen(str string) int {

	nub := 0
	for _, i := range str {

		if i == i {
			nub++
		}

	}
	return nub

}
