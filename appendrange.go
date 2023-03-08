package piscine

func AppendRange(min, max int) []int {
	mas := []int{}

	if max > min {
		for i := min; i < max; i++ {
			mas = append(mas, i)
		}
		return mas
	} else {
		return nil
	}
}
