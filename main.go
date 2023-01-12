package main

import "fmt"

func main() {
	// Go is statically typed
	var card string = newCard()
	// card := "Ace of Spades"
	// := is used for initialization  
	fmt.Println(card)
}

func newCard() string {
	return "Ace of Spades" 
}