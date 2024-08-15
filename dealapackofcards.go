package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	for _, v := range deck {
		switch v {
		case 1:
			fmt.Printf("Player 1: %d, %d, %d\n", deck[0], deck[1], deck[2])
		case 3:
			fmt.Printf("Player 2: %d, %d, %d\n", deck[3], deck[4], deck[5])
		case 6:
			fmt.Printf("Player 3: %d, %d, %d\n", deck[6], deck[7], deck[8])
		case 9:
			fmt.Printf("Player 4: %d, %d, %d\n", deck[9], deck[10], deck[11])
		}
	}
}
