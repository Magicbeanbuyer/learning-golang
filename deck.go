package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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
	values := [4]string{"a", "1", "2", "3"}

	newDeck := deck{}

	for _, suit := range suits {
		for _, value := range values {
			newDeck = append(newDeck, value+" of "+suit)
		}
	}

	return newDeck
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	err := os.WriteFile(filename, []byte(d.toString()), 0666)
	return err
}

func loadDeckFromFile(filename string) deck {
	byteSlice, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	deckString := string(byteSlice)
	deckSlick := strings.Split(deckString, ",")
	return deckSlick
}

func (d deck) shuffle() {
	deckLength := len(d)
	seed := time.Now().UnixNano()
	mySource := rand.NewSource(seed)
	myRandomGenerator := rand.New(mySource)

	for index := range d {
		randomNumber := myRandomGenerator.Intn(deckLength - 1)
		d[index], d[randomNumber] = d[randomNumber], d[index]
	}
}

/*
Go does not have class, declare custom type instead

string type var cannot have nil as its value

when a var is declared but not initialized, so-called "zero value" (0, "", false)
is assigned to it.

go variable uses camelCase

:= is for declaration + assignment, whereas = is for assignment only.
*/
