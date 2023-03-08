package piscine

func IterativeFactorial(nb int) int {
	if nb == 1 || nb == 0 {
		return 1
	} else if nb < 0 || nb > 20 {
		return 0
	} else {
		return nb * (nb - 1)
	}
}
