/////////. MAIN.GO //////////////



package main

func main() {
	cards := newDeck()

	cards.print()
}





//////////// DECK.GO /////////



package main

import "fmt"

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
