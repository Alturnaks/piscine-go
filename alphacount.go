package piscine

func AlphaCount(s string) int {
	count := 0
	mas := []rune(s)

	for _, j := range mas {
		if j >= 'a' && j <= 'z' || j >= 'A' && j <= 'Z' {
			count++
		}
	}
	return count
}
