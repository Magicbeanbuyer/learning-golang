package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingDeck := cards.deal(5)
	fmt.Println(hand.toString())
	remainingDeck.print()
	err := hand.saveToFile("hand.txt")
	print(err)
}

/* NOTES
Go is statically typed

:= is used for initialization

array: fixed length, elements must have identical type
slice: varied length, elements must have identical type

append does not modify the slice in place, rather return a new slice
*/
