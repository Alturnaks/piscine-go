package piscine

import (
	"github.com/01-edu/z01"
)

func splitnumb(n int) {
	k := 0
	var mas []int
	for n != 0 {
		k = n % 10
		if k < 0 {
			mas = append(mas, k*-1)
		} else {
			mas = append(mas, k)
		}

		n = n / 10

	}

	for i := len(mas) - 1; i >= 0; i-- {
		z01.PrintRune(rune(mas[i] + 48))
	}
}

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else if n < 0 {
		z01.PrintRune('-')
		splitnumb(n)
	} else {
		splitnumb(n)
	}
}
