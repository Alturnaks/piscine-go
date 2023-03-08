package piscine

func JumpOver(str string) string {
	if str == "" {
		return "\n"
	}
	res := ""
	for i := 0; i < len(str); i++ {
		if (i+1)%3 == 0 {
			res += string(str[i])
		}
	}
	res += "\n"
	return res
}
