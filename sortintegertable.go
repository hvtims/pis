package piscine

func SortIntegerTable(table []int) {
	n := len(table)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if table[j] < table[minIndex] {
				minIndex = j
			}
		}
		table[i], table[minIndex] = table[minIndex], table[i]
	}
}
