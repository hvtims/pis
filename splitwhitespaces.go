package piscine

func SplitWhiteSpaces(s string) []string {
	str := ""
	table := []string{}
	isSpaced := false
	for i, v := range s {
		if v == ' ' && !isSpaced {
			table = append(table, str)
			str = ""
			isSpaced = true
		} else if v != ' ' {
			str += string(v)
			isSpaced = false
		}
		if i == len(s)-1 {
			table = append(table, str)
		}
	}
	return table
}
