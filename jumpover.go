package piscine

func JumpOver(str string) string {
	result := ""

	if len(str) < 2 || str == "" {
		return "\n"
	}
	for i := 1; i <= len(str); i++ {
		if i%3 == 0 {
			result += string(str[i-1])
		}
	}
	return result + "\n"
}
