package piscine

func ActiveBits(n int) int {
	result := 0
	for n > 0 {
		if n&1 == 1 {
			result++
		}
		n >>= 1
	}
	return result
}
