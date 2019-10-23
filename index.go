package piscine

func Index(s string, toFind string) int {

	str1 := []rune(s)
	str2 := []rune(toFind)
	index := 0
	u := false
	len1 := 0
	len2 := 0
	for range str1 {
		len1++
	}
	for range str2 {
		len2++
	}
	if len1 < len2 {
		return -1
	}
	if len1 <= 0 || len2 <= 0 {
		return 0
	}
	for i1, letter1 := range str1 {
		if str2[0] == letter1 {
			index = i1
			for i2, letter2 := range str2 {
				if str1[index+i2] != letter2 {
					u = false
					break

				} else {
					u = true

				}
			}

			if u {
				return index
			}
		}
	}

	return -1

}
