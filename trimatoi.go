package piscine

func TrimAtoi(s string) int {
	num := 0
	res := 1
	for _, i := range s {
		if i >= '0' && i <= '9' {
			b := int(i - 48)
			num = num*10 + b
		} else if num == 0 && i == '-' {
			res = res * (-1)
		}
	}
	return num * res
}
