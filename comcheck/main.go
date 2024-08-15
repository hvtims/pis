package main

import (
	"fmt"
	"os"
)

func main() {
	for i := range os.Args {
		if os.Args[i] == "galaxy" || os.Args[i] == "01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
