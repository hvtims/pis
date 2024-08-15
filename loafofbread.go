package piscine

func LoafOfBread(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	var result string
	j := 0
	for i := 0; i < len(str); i++ {
		if j < 5 && str[i] == ' ' {
			continue
		}
		if j == 5 {
			if i != len(str)-1 && str[i+1] == ' ' {
				continue
			}
			if i == len(str)-1 {
				break
			}
			result += " "
			j = 0
			continue
		}
		result += string(str[i])
		j++
	}
	result += "\n"
	return result
}
