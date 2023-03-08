package piscine

func ConcatParams(args []string) string {
	str := ""
	for i, j := range args {
		if i == 0 {
			str += j
		} else {
			j = "\n" + j
			str = str + j
		}
	}
	return str
}
