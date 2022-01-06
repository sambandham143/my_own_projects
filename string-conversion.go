###############.  MAIN.GO ################

package main

import "fmt"

func main() {
	cards := newDeck()

	fmt.Println(cards.toString())

}




#################. DECK.GO #############



package main

import (
	"fmt"
	"strings"
)

type deck []string

////// This is for creating a New Deck

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}

//////// This is for Printing

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

///////// This is a deal function

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

/////// THis is string conversion function

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
