package piscine

func StringToIntSlice(str string) []int {
	A := []rune(str)
	var B []int
	for i := 0; i < len(A); i++ {
		B = append(B, int(A[i]))
	}
	return B
}
