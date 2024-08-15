package piscine

func IsLower(s string) bool {
	for _, sa := range s {
		if sa < 'a' || sa > 'z' {
			return false
		}
	}
	return true
}
