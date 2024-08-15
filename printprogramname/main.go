package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument1 := []rune(os.Args[0])

	if len(os.Args) > 0 {
		for i := range argument1 {
			if argument1[i] != '/' && argument1[i] != '.' {
				z01.PrintRune(argument1[i])
			}
		}
		z01.PrintRune('\n')
	}
}
