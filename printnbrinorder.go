package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	k := 0
	var mas []int
	for n != 0 {
		k = n % 10

		mas = append(mas, k)

		n = n / 10
	}
	for j := 0; j <= 9; j++ {
		for i := 0; i < len(mas); i++ {
			if j == mas[i] {
				z01.PrintRune(rune(mas[i] + 48))
			}
		}
	}
}
