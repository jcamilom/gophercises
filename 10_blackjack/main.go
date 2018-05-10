package main

import (
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

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == deck.Ace {
			// ace is currently worth 1, and we are changing it to
			// be worth 11, 11 - 1 = 10
			return minScore + 10
		}
	}
	return minScore
}

func (h Hand) MinScore() int {
	score := 0
	for _, c := range h {
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	//theDeck := deck.New(deck.Deck(3), deck.Shuffle)
	var gs GameState
	gs.Deck = deck.New(deck.Deck(3), deck.Shuffle)
	/* var card deck.Card
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
	// If the dealer score <= 16, we hit.
	// If the dealer has a soft 17, then we hit.
	for dealerH.Score() <= 16 || (dealerH.Score() == 17 && dealerH.MinScore() != 17) {
		card, theDeck = draw(theDeck)
		dealerH = append(dealerH, card)
	}
	pScore, dScore := playerH.Score(), dealerH.Score()
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player: ", playerH, "\nScore:", pScore)
	fmt.Println("Dealer: ", dealerH, "\nScore:", dScore)
	switch {
	case pScore > 21:
		fmt.Println("You busted")
	case dScore > 21:
		fmt.Println("Dealer busted")
	case pScore > dScore:
		fmt.Println("You win!")
	case dScore > pScore:
		fmt.Println("You lose")
	case dScore == pScore:
		fmt.Println("Draw")
	} */

}

func draw(d []deck.Card) (deck.Card, []deck.Card) {
	return d[0], d[1:]
}

type State int8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

type GameState struct {
	Deck   []deck.Card
	Turn   State
	Player Hand
	Dealer Hand
}

func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.Turn {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Player
	default:
		panic("it isn't currently any player's turn")
	}
}

func clone(gs GameState) GameState {
	ret := GameState{
		Deck:   make([]deck.Card, len(gs.Deck)),
		Turn:   gs.Turn,
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
	}
	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Player)
	copy(ret.Dealer, gs.Dealer)
	return ret
}
