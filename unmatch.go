package piscine

func Unmatch(a []int) int {
	d := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] {
				d++
			}
		}
		if d%2 == 1 {
			return a[i]
		}
	}
	return -1
}
