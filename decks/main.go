package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingDeck := cards.deal(5)
	fmt.Println(hand.toString())
	remainingDeck.print()
	fileName := "deck.txt"
	err := cards.saveToFile(fileName)
	fmt.Println(err)
	savedDeck := loadDeckFromFile(fileName)
	fmt.Println(savedDeck)
	//loadDeckFromFile("fakename")
	cards.shuffle()
	cards.print()

}

/* NOTES
Go is statically typed

:= is used for initialization

x := 42

var x int = 42

var x int
x = 42


array: fixed length, elements must have identical type
slice: varied length, elements must have identical type

append does not modify the slice in place, rather return a new slice

go test files have suffix _test, e.g. deck_test.go
*/
