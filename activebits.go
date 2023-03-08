package piscine

func ActiveBits(n int) int {
	s := 0
	for n > 0 {
		if n&1 == 1 {
			s++
		}
		n >>= 1
	}
	return s
}
