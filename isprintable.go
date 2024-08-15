package piscine

func IsPrintable(s string) bool {
	for _, la := range s {
		if !(la >= ' ' && la <= '~') {
			return false
		}
	}
	return true
}
