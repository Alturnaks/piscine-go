package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args

	for i := len(a) - 1; i >= 1; i-- {
		f := []rune(a[i])
		for j := range a[i] {
			z01.PrintRune(f[j])
		}
		z01.PrintRune('\n')
	}
}
