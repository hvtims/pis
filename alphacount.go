package piscine

func AlphaCount(s string) int {
	count := 0

	for _, la := range s {
		if (la >= 'a' && la <= 'z') || (la >= 'A' && la <= 'Z') {
			count++
		}
	}

	return count
}
