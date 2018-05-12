package blackjack

import (
	"fmt"

	"github.com/jcamilom/gophercises/09_deck"
)

type AI interface {
	Bet() int
	Results(hand []deck.Card, dealer []deck.Card)
	Play(hand []deck.Card, dealer deck.Card) Move
}

type HumanAI struct{}

func (ar *HumanAI) Play(hand []deck.Card, dealer deck.Card) Move {
	for {
		fmt.Println("Player: ", hand)
		fmt.Println("Dealer: ", dealer)
		fmt.Println("What will you do? (h)it, (s)tand")
		var input string
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return Hit
		case "s":
			return Stand
		default:
			fmt.Println("Invalid option:", input)
		}
	}
}

func (ai *HumanAI) Bet() int {
	return 1
}

func (ai *HumanAI) Results(hand [][]deck.Card, dealer []deck.Card) {
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player: ", hand)
	fmt.Println("Dealer: ", dealer)
}

// Filler to be implemented

type Move func(GameState) GameState

func Hit(gs GameState) GameState {
	return gs
}

func Stand(gs GameState) GameState {
	return gs
}
