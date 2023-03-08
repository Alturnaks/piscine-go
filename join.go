package piscine

func Join(strs []string, sep string) string {
	str := ""
	for i, j := range strs {
		if i == len(strs)-1 {
			str += j
			break
		}
		str = str + j + sep
	}
	return str
}
