package piscine

func StrLen(s string) int {
	var count int
	mas := []rune(s)

	for i := range mas {
		count = i
	}

	return count + 1
}
