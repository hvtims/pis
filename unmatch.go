package piscine

func Unmatch(a []int) int {
	for _, pa := range a {
		count := 0
		for _, v := range a {
			if v == pa {
				count++
			}
		}
		if count%2 != 0 {
			return pa
		}
	}
	return -1
}
