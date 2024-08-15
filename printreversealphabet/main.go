package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for ii := 122; ii >= 97; ii-- {
		z01.PrintRune(rune(ii))
	}
	z01.PrintRune('\n')
}
