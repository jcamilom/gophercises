package main

import (
	"fmt"

	"github.com/jcamilom/gophercises/11_blackjacj_ai/blackjack"
)

func main() {
	game := blackjack.New()
	winnings := game.Play(blackjack.HumanAI())
	fmt.Println(winnings)
}
