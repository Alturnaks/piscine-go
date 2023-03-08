package piscine

func ToLower(s string) string {
	mas := []rune(s)
	for i := 0; i < len(mas); i++ {
		if mas[i] >= 65 && mas[i] <= 90 {
			mas[i] = rune(mas[i] + 32)
		}
	}
	return string(mas)
}
