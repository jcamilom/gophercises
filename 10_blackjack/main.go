package main

import (
	"fmt"
	"strings"

	"github.com/jcamilom/gophercises/09_deck"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) DealerString() string {
	return h[0].String() + ", **HIDDEN**"
}

func main() {
	theDeck := deck.New(deck.Deck(3), deck.Shuffle)
	var card deck.Card
	var playerH, dealerH Hand
	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&playerH, &dealerH} {
			card, theDeck = draw(theDeck)
			*hand = append(*hand, card)
		}
	}
	var input string
	for input != "s" {
		fmt.Println("Player: ", playerH)
		fmt.Println("Dealer: ", dealerH.DealerString())
		fmt.Println("What will you do? (h)it, (s)tand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			card, theDeck = draw(theDeck)
			playerH = append(playerH, card)
		}
	}
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player: ", playerH)
	fmt.Println("Dealer: ", dealerH)
}

func draw(d []deck.Card) (deck.Card, []deck.Card) {
	return d[0], d[1:]
}
