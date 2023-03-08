package piscine

func IsNumeric(s string) bool {
	mas := []rune(s)
	for i := 0; i < len(mas); i++ {
		if !(mas[i] >= '0' && mas[i] <= '9') {
			return false
		}
	}
	return true
}
