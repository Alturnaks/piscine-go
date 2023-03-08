package piscine

func Index(s string, toFind string) int {
	mas := []rune(s)
	if s == toFind {
		return 0
	} else {
		for i := 0; i < len(s); i++ {
			if string(mas[i:len(toFind)+i]) == toFind {
				return i
			}
		}
		return -1
	}
}
