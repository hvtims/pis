package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 0 {
		for argsIndex := len(os.Args) - 1; argsIndex >= 1; argsIndex-- {
			argTable := []rune(os.Args[argsIndex])
			for index := range argTable {
				z01.PrintRune(argTable[index])
			}
			z01.PrintRune('\n')
		}
	}
}
