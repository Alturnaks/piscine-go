package piscine

func StrRev(s string) string {
	a := ""
	for i := len(s) - 1; i >= 0; i-- {
		a = a + string(s[i])
	}
	return a
}
