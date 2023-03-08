package piscine

func Capitalize(s string) string {
	mas := []rune(s)
	for i, j := range mas {
		if i == 0 || !check(mas[i-1]) {
			if j >= 'a' && j <= 'z' {
				mas[i] = j - 32
			}
		} else {
			if j >= 'A' && j <= 'Z' {
				mas[i] = j + 32
			}
		}
	}
	return string(mas)
}

func check(a rune) bool {
	if a >= 'a' && a <= 'z' {
		return true
	}
	if a >= 'A' && a <= 'Z' {
		return true
	}
	if a >= '0' && a <= '9' {
		return true
	}
	return false
}
