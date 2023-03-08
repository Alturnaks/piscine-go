package piscine

func IsPrintable(s string) bool {
	mas := []rune(s)
	for i := 0; i < len(mas); i++ {
		if !(mas[i] >= 32 && mas[i] <= 126) {
			return false
		}
	}
	return true
}
