package piscine

func BasicAtoi2(s string) int {
	result := 0
	for _, digit := range s {
		if digit < '0' || digit > '9' {
			return 0
		}
		digitValue := int(digit - '0')
		result = result*10 + digitValue

	}
	return result
}
