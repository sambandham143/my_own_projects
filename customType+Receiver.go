////// MAIN.GO ///////


package main

func main() {
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}




/////// DECK.GO ///////


package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
