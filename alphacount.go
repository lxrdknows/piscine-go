package piscine


func AlphaCount(str string) int {
	counter := 0
	
	for _, count := range str {
		if count >= 'A' && count <= 'Z' {

			counter++
		} else if count >= 'a' && count <= 'z' {
			counter++

		}
		
		}
		return counter
	}
	

}

