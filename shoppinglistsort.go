package piscine

func ShoppingListSort(slice []string) []string {
	var qq string
	for i := 0; i < len(slice); i++ {
		for j := i; j < len(slice); j++ {
			if len(slice[i]) > len(slice[j]) {
				qq = slice[i]
				slice[i] = slice[j]
				slice[j] = qq
			}
		}
	}
	return slice
}
