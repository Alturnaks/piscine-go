package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args
	for i := 0; i < len(a)-1; i++ {
		for j := 1; j < len(a)-i-1; j++ {
			if a[j] < a[j+1] {
				t := a[j]
				a[j] = a[j+1]
				a[j+1] = t
			}
		}
	}
	for i := len(a) - 1; i >= 1; i-- {
		f := []rune(a[i])
		for j := range a[i] {
			z01.PrintRune(f[j])
		}
		z01.PrintRune('\n')
	}
}
