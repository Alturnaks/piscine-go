package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func dv(n int) {
	k := n % 10
	z01.PrintRune(rune((n-(k%10))/10) + 48)
	z01.PrintRune(rune(k + 48))
}

func main() {
	points := &point{}
	setPoint(points)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	dv(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')

	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	dv(points.y)
	z01.PrintRune('\n')
}
