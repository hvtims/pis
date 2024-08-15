package piscine

func Join(strs []string, sep string) string {
	newString := ""

	for i := range strs {
		if i != len(strs)-1 {
			newString += strs[i] + sep
		} else {
			newString += strs[i]
		}
	}

	return newString
}
