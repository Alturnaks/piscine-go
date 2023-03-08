package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 4 {
		var result int
		Rn_1 := []rune(os.Args[1])
		Rn_2 := []rune(os.Args[3])
		var val_1 int
		var val_2 int
		for i := 0; i < len(Rn_1); i++ {
			for j := 0; j <= 9; j++ {
				if Rn_1[i] == rune(j+48) {
					val_1 = 10*val_1 + j
				}
			}
		}
		for i := 0; i < len(Rn_2); i++ {
			for j := 0; j <= 9; j++ {
				if Rn_2[i] == rune(j+48) {
					val_2 = 10*val_2 + j
				}
			}
		}
		if Rn_1[0] == '-' {
			val_1 *= (-1)
		}
		if Rn_2[0] == '-' {
			val_2 *= (-1)
		}
		errorMsg := "No division by "
		errorMsg2 := "No modulo by "
		if os.Args[2] == "*" {
			result = val_1 * val_2
		} else if os.Args[2] == "/" {
			if val_2 == 0 {
				for i := 0; i < len(errorMsg); i++ {
					for j := 0; j <= 127; j++ {
						if rune(errorMsg[i]) == rune(j) {
							z01.PrintRune(rune(j))
						}
					}
				}
			} else {
				result = val_1 / val_2
			}
		} else if os.Args[2] == "+" {
			result = val_1 + val_2
		} else if os.Args[2] == "-" {
			result = val_1 - val_2
		} else if os.Args[2] == "%" {
			if val_2 == 0 {
				for i := 0; i < len(errorMsg2); i++ {
					for j := 0; j <= 127; j++ {
						if rune(errorMsg2[i]) == rune(j) {
							z01.PrintRune(rune(j))
						}
					}
				}
			} else {
				result = val_1 % val_2
			}
		}
		count_1 := 0
		count_2 := 0
		for i := 0; i < len(Rn_1); i++ {
			if Rn_1[i] >= '0' && Rn_1[i] <= '9' || Rn_1[i] <= '-' {
				count_1++
			}
		}
		for i := 0; i < len(Rn_2); i++ {
			if Rn_2[i] >= '0' && Rn_2[i] <= '9' || Rn_1[i] <= '-' {
				count_2++
			}
		}
		op := []rune("+-*/%")
		operator := []rune(os.Args[2])
		count_op := 0
		for i := 0; i < len(op); i++ {
			if op[i] == operator[0] {
				count_op++
			}
		}
		if count_1 == len(Rn_1) && count_2 == len(Rn_2) && count_op == 1 && len(Rn_1) <= 18 && len(Rn_2) <= 18 {
			if result < 0 {
				result *= (-1)
				var A []int
				k := 0
				for result != 0 {
					k = result % 10
					A = append(A, k)
					result /= 10
				}
				z01.PrintRune('-')
				for i := len(A) - 1; i >= 0; i-- {
					z01.PrintRune(rune(A[i] + 48))
				}
			} else if result == 0 {
				z01.PrintRune('0')
			} else {
				var A []int
				k := 0
				for result != 0 {
					k = result % 10
					A = append(A, k)
					result /= 10
				}
				for i := len(A) - 1; i >= 0; i-- {
					z01.PrintRune(rune(A[i] + 48))
				}
			}
			z01.PrintRune('\n')
		}
	}
}
