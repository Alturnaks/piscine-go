package piscine

func NRune(s string, n int) rune {
	r := []rune(s)
	if n > len(s) {
		return 0
	} else if n > 0 {
		return rune(r[n-1])
	}
	return 0
}
