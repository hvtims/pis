package piscine

func DescendAppendRange(max, min int) []int {
	if max <= min {
		return []int{}
	}

	newArray := []int{}

	for i := max; i > min; i-- {
		newArray = append(newArray, i)
	}
	return newArray
}
