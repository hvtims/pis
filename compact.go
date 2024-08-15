package piscine

func Compact(ptr *[]string) int {
	myslice := *ptr
	myslice2 := []string{}

	for i := 0; i < len(myslice); i++ {
		if myslice[i] != "" {
			myslice2 = append(myslice2, myslice[i])
		}
	}
	*ptr = myslice2

	return len(myslice2)
}
