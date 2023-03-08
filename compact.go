package piscine

func Compact(ptr *[]string) int {
	var str []string
	for i := 0; i < len(*ptr); i++ {
		if len((*ptr)[i]) != 0 {
			str = append(str, (*ptr)[i])
		}
	}
	*ptr = str
	return len((*ptr))
}
