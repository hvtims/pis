package piscine

func Map(f func(int) bool, a []int) (result []bool) {
	for _, la := range a {
		result = append(result, f(la))
	}
	return
}
