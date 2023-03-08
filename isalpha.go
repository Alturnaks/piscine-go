package piscine

func IsAlpha(s string) bool {
	mas := []rune(s)
	for i := 0; i < len(mas); i++ {
		if !(mas[i] >= 'a' && mas[i] <= 'z' || mas[i] >= 'A' && mas[i] <= 'Z' || mas[i] >= '0' && mas[i] <= '9') {
			return false
		}
	}
	return true
}
