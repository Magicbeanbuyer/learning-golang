package main

import (
	"fmt"
	"strings"
)

type deck []string

// d: function parameter, use type abbreviation by convention, deck: parameter type
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	suits := [4]string{"spades", "hearts", "diamonds", "clubs"}
	values := [13]string{"a", "1", "2", "3", "4", "5", "6", "7", "8", "10", "j", "q", "k"}

	newDeck := deck{}

	for _, suit := range suits {
		for _, value := range values {
			newDeck = append(newDeck, value+" of "+suit)
		}
	}

	return newDeck
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func toString(d deck) string {
	return strings.Join(d, ",")
}

/*
Go does not have class, declare custom type instead

string type var cannot have nil as its value

when a var is declared but not initialized, so-called "zero value" (0, "", false)
is assigned to it.
*/
