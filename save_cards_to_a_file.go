############### MAIN.GO #################


package main

func main() {
	cards := newDeck()

	cards.saveToFile("my_cards")

}



###############. DECK.GO ##################


package main

import (
	"fmt"
	"io/ioutil"
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

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}
