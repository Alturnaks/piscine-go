package piscine

func Split(s, sep string) []string {
	var result []string
	var oneStr []rune
	runeS := []rune(s)
	runeSep := []rune(sep)
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if runeS[j] == runeSep[0] && s[j:(j+len(sep))] == sep {
				break
			} else {
				oneStr = append(oneStr, runeS[j])
			}
			i = j + len(sep)
		}

		if len(oneStr) > 0 && oneStr[0] != ' ' {
			result = append(result, string(oneStr))
		}
		oneStr = []rune{}
	}
	return result
}
