package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)
	//hand.print()
	fmt.Println(toString(hand))
	remainingDeck.print()
}

/* NOTES
Go is statically typed

:= is used for initialization

array: fixed length, elements must have identical type
slice: varied length, elements must have identical type

append does not modify the slice in place, rather return a new slice
*/
