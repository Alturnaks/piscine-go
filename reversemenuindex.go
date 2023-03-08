package piscine

func ReverseMenuIndex(menu []string) []string {
	mas := make([]string, len(menu))
	for i := 0; i < len(mas); i++ {
		mas[i] = menu[len(menu)-1-i]
	}
	return mas
}
