package piscine

// import "fmt"

func ShoppingSummaryCounter(str string) map[string]int {
	res := make(map[string]int)
	word := ""
	if str == "" {
		res[word] = 1
		return res
	} else if str == " " {
		res[word] = 2
		return res
	}
	count := 0

	for j, i := range str {
		if i == ' ' {
			v := res[word]
			res[word] = v + 1
			word = ""
			count++
		} else if j == len(str)-1 {
			word += string(i)
			v := res[word]
			res[word] = v + 1
			word = ""
		} else {
			word += string(i)
		}
	}
	if count == 1 {
		res[word] = 1
	}
	return res
}

// func main() {
// 	summary := "#$sds "
// 	for index, element := range ShoppingSummaryCounter(summary) {
// 		fmt.Println(index, "=>", element)
// 	}
// }
