package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 0 {
		for argsIndex := 1; argsIndex < len(os.Args); argsIndex++ {
			argTable := []rune(os.Args[argsIndex])
			for index := range argTable {
				z01.PrintRune(argTable[index])
			}
			z01.PrintRune('\n')
		}
	}
}
