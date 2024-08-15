package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascendingOrdered := true
	descendingOrdered := true

	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			descendingOrdered = false
		}
	}

	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			ascendingOrdered = false
		}
	}

	return ascendingOrdered || descendingOrdered
}
