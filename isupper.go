package piscine

func IsUpper(s string) bool {
	mas := []rune(s)
	for i := 0; i < len(mas); i++ {
		if mas[i] < 'A' || mas[i] > 'Z' {
			return false
		}
	}
	return true
}
