package piscine

func Map(f func(int) bool, a []int) []bool {
	mas := []bool{}
	for _, i := range a {
		mas = append(mas, f(i))
	}
	return mas
}
