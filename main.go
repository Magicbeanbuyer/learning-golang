package main

func main() {
	cards := newDeck()
	cards.print()
}

/* NOTES
Go is statically typed

:= is used for initialization 

array: fixed length, elements must have identical type
slice: varied length, elements must have identical type

append does not modify the slice in place, rather return a new slice
*/