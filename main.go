package main

import (
	"fmt"
)

func main() {
	cards := []string{newCard(), "Five of Diamonds"}
	cards = append(cards, "Six of Spades")

	for i, card := range cards{
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Spades" 
}

/* NOTES
Go is statically typed

:= is used for initialization 

array: fixed length, elements must have identical type
slice: varied length, elements must have identical type

append does not modify the slice in place, rather return a new slice
*/