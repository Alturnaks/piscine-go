package piscine

func DescendAppendRange(max, min int) []int {
	var A []int
	if max > min {
		for i := max; i > min; i-- {
			A = append(A, i)
		}
	} else {
		return []int{}
	}
	return A
}
