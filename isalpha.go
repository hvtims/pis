package piscine

func IsAlpha(s string) bool {
	for _, la := range s {
		if !((la >= 'a' && la <= 'z') || (la >= 'A' && la <= 'Z') || (la >= '0' && la <= '9')) {
			return false
		}
	}
	return true
}
