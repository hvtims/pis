package piscine

func TrimAtoi(s string) int {
	var numStr string
	isNegative := false
	signFound := false

	for _, ch := range s {
		if ch == '-' && !signFound {
			isNegative = true
			signFound = true
		} else if ch == '+' && !signFound {
			signFound = true
		} else if ch >= '0' && ch <= '9' {
			numStr += string(ch)
			signFound = true
		}
	}

	if numStr == "" {
		return 0
	}

	num := 0
	for _, ch := range numStr {
		digit := int(ch - '0')
		num = num*10 + digit
	}

	if isNegative {
		num = -num
	}

	return num
}
