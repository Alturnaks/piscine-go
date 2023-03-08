package piscine

func ToUpper(s string) string {
	mas := []rune(s)
	for i := 0; i < len(mas); i++ {
		if mas[i] >= 97 && mas[i] <= 122 {
			mas[i] = rune(mas[i] - 32)
		}
	}
	return string(mas)
}
