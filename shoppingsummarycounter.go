package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)

	word := ""

	for i := 0; i < len(str); i++ {
		if string(str[i]) == " " {
			result[word]++
			word = ""
		} else {
			word += string(str[i])
		}
	}
	result[word]++

	return result
}
