package piscine

func Rot14(s string) string {
	runeS := []rune(s)
	for i := 0; i < len(s); i++ {
		if runeS[i] >= 'a' && runeS[i] <= 'z'-14 {
			runeS[i] = runeS[i] + 14
		} else if runeS[i] > 'z'-14 && runeS[i] <= 'z' {
			runeS[i] = 'a' + 13 + runeS[i] - 'z'
		} else if runeS[i] >= 'A' && runeS[i] <= 'Z'-14 {
			runeS[i] = runeS[i] + 14
		} else if runeS[i] > 'Z'-14 && runeS[i] <= 'Z' {
			runeS[i] = 'A' + 13 + runeS[i] - 'Z'
		}
	}
	return string(runeS)
}
