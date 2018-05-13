package main

import (
	"fmt"

	deck "github.com/jcamilom/gophercises/09_deck"
	"github.com/jcamilom/gophercises/11_blackjacj_ai/blackjack"
)

type basicAI struct{}

func (ai *basicAI) Bet(shuffled bool) int {
	panic("not implemented")
}

func (ai *basicAI) Play(hand []deck.Card, dealer deck.Card) blackjack.Move {
	panic("not implemented")
}

func (ai *basicAI) Results(hands [][]deck.Card, dealer []deck.Card) {
	panic("not implemented")
}

func main() {
	opts := blackjack.Options{
		Decks:           3,
		Hands:           2,
		BlackjackPayout: 1.5,
	}
	game := blackjack.New(opts)
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
