package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	count := 0
	count_2 := 0
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) >= 0 {
			count++
		} else if f(a[i-1], a[i]) <= 0 {
			count_2++
		}
	}
	if count == len(a)-1 || count_2 == len(a)-1 {
		return true
	} else {
		return false
	}
}

func IsEqual(a, b int) int {
	return a - b
}
