package piscine

func Sqrt(value int) int {
	if value < 0 {
		return 0
	}

	for i := 0; i*i <= value; i++ {
		if i*i == value {
			return i
		}
	}
	return 0
}
