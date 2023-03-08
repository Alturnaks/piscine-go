package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	c := 1
	for j := 3; j < 13; j += 3 {

		fmt.Printf("Player %d: %d, %d, %d\n", c, deck[j-3], deck[j-2], deck[j-1])
		c++
	}
}
