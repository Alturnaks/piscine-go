package piscine

import "github.com/01-edu/z01"

func outputDigit(f, s int) {
	z01.PrintRune(rune(f + 48))
	z01.PrintRune(rune(s + 48))
}

func PrintComb2() {
	for i := 0; i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			outputDigit((i-(i%10))/10, (i % 10))
			z01.PrintRune(rune(' '))
			outputDigit((j-(j%10))/10, (j % 10))
			if i == 98 && j == 99 {
				break
			}
			z01.PrintRune(',')
			z01.PrintRune(' ')

		}
	}
	z01.PrintRune('\n')
}
