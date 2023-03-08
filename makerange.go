package piscine

func MakeRange(min, max int) []int {
	if max > min {
		mas := make([]int, max-min)
		for i := 0; i < max-min; i++ {
			mas[i] = i + min
		}
		return mas
	} else {
		return nil
	}
}
