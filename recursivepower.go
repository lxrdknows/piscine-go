package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else if power == 1 {
		return nb
	}
	res := 1
	res *= nb
	return res * RecursivePower(res, power-1)
}
