package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	result := args[0]
	for _, arg := range args[1:] {
		result += "\n" + arg
	}

	return result
}
