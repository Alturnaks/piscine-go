package piscine

func LastRune(s string) rune {
	r := []rune(s)
	return rune(r[len(r)-1])
}
