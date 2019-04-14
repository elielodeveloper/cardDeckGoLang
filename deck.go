package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//Create a new type of 'deck'
//which is a slice of strings
type deck []string

//It is a kind of constructor
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Curinga"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//Print the of the deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//Create hands
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//Change a deck to a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
