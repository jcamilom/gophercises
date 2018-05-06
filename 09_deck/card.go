//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Jocker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Jocker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New() []Card {
	theDeck := make([]Card, 0, ((len(suits)) * int((maxRank))))
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			theDeck = append(theDeck, Card{Suit: suit, Rank: rank})
		}
	}
	return theDeck
}
