############  MAIN.GO ############


package main

import "fmt"

func main() {
	cards := newDeck()

	fmt.Println("Cards with original order")

	cards.print()

	fmt.Println("Cards with after shuffle order")

	cards.shuffle()

	cards.print()

}






################. DECK.GO ###############



package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

//// This is saving a file to a machine function

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

///// Reading a file from a machine function

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

////////  Shuffle the Deck function

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
