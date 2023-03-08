package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args[0]
	s := []rune(a)
	for i := 2; i < len(s); i++ {
		z01.PrintRune(s[i])
	}
	z01.PrintRune('\n')
}
