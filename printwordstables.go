package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			runeA := a[i]
			z01.PrintRune(rune(runeA[j]))
		}
		z01.PrintRune('\n')
	}
}
