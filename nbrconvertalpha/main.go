package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	flag := false
	start := 1
	if len(args) < 2 {
		return
	}
	if args[1] == "--upper" {
		flag = true
		start = 2
	}
	for i := start; i < len(args); i++ {
		nb := letter(args[i])
		if nb < 1 || nb > 26 {
			z01.PrintRune(' ')
		} else if !flag {
			nb += 96
			z01.PrintRune(rune(nb))
		} else if flag {
			nb += 64
			z01.PrintRune(rune(nb))
		}
	}
	z01.PrintRune('\n')
}

func letter(s string) int {
	runes := []rune(s)
	count := 0
	sign := '+'
	res := 0

	for i := 0; i < len(runes); i++ {
		if (runes[i] < 48 || runes[i] > 57) && res != 0 {
			return 0
		}
		if (runes[i] == '+' || runes[i] == '-') && res == 0 {
			if runes[i] == '-' {
				sign = '-'
			}
			count++
		}
		if runes[i] == ' ' || (runes[i] >= 65 && runes[i] <= 122) {
			return 0
		}
		if count > 1 {
			return 0
		}
		if sign == '-' {
			if runes[i] >= '0' && runes[i] <= '9' {
				res = res*10 - int(runes[i]-'0')
			}
		} else if sign == '+' {
			if runes[i] >= '0' && runes[i] <= '9' {
				res = res*10 + int(runes[i]-'0')
			}
		} else {
			return 0
		}
	}
	return res
}
