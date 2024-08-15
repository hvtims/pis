package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for ii := '0'; ii <= '9'; ii++ {
		z01.PrintRune(rune(ii))
	}
	z01.PrintRune('\n')
}
