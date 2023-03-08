package piscine

func Abort(a, b, c, d, e int) int {
	primes := []int{a, b, c, d, e}
	s := 0
	for i := 0; i < len(primes); i++ {
		for j := 0; j < len(primes)-1; j++ {
			if primes[j] > primes[j+1] {
				s = primes[j]
				primes[j] = primes[j+1]
				primes[j+1] = s
			}
		}
	}
	return primes[2]
}
