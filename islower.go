package piscine

func IsLower(s string) bool {
	mas := []rune(s)
	for i := 0; i < len(mas); i++ {
		if mas[i] < 'a' || mas[i] > 'z' {
			return false
		}
	}
	return true
}
