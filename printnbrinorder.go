package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var count [10]int
	for n > 0 {
		count[n%10]++
		n /= 10
	}

	for i, v := range count {
		for ; v > 0; v-- {
			z01.PrintRune(rune('0' + i))
		}
	}
}
