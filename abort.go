package piscine

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func Abort(a, b, c, d, e int) int {
	arg := []int{a, b, c, d, e}
	bubbleSort(arg)
	return arg[2]
}
