package piscine

func BasicAtoi(s string) int {
	result := 0
	for _, digit := range s {
		digitValue := int(digit - '0')
		result = result*10 + digitValue

	}
	return result
}
