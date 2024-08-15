package piscine

func IsUpper(s string) bool {
	for _, sa := range s {
		if sa < 'A' || sa > 'Z' {
			return false
		}
	}
	return true
}
