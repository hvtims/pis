package piscine

func ShoppingListSort(slice []string) []string {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice); j++ {
			a := slice[i]
			if len(slice[i]) > len(slice[i+1]) {
				slice[i] = slice[i+1]
				slice[i+1] = a
			}
		}
	}
	return slice
}
