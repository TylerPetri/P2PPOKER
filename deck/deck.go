package deck

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Suit int

func (s Suit) String() string {
	switch s {
	case Spades:
		return "SPADES"
	case Hearts:
		return "HEARTS"
	case Diamonds:
		return "DIAMONDS"
	case Clubs:
		return "CLUBS"
	default:
		panic("invalid card suit")
	}
}

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

type Card struct {
	suit  Suit
	value int
}

func (c Card) String() string {
	value := strconv.Itoa(c.value)
	if c.value == 1 {
		value = "ACE"
	}
	return fmt.Sprintf("%s of %s %s", value, c.suit, suitToUnicode(c.suit))
}

func NewCard(s Suit, v int) Card {
	if v > 13 {
		panic("the value of the card cannot be higher than 13")
	}
	return Card{
		suit:  s,
		value: v,
	}
}

type Deck [52]Card

func New() Deck {
	var (
		nSuits = 4
		nCards = 13
		d      = [52]Card{}
	)

	x := 0
	for i := range nSuits {
		for j := range nCards {
			d[x] = NewCard(Suit(i), j+1)
			x++
		}
	}

	return shuffle(d)
}

func shuffle(d Deck) Deck {
	for i := range len(d) {
		r := rand.Intn(i + 1)

		if r != i {
			d[i], d[r] = d[r], d[i]
		}
	}

	return d
}

func suitToUnicode(s Suit) string {
	switch s {
	case Spades:
		return "♠"
	case Hearts:
		return "♥"
	case Diamonds:
		return "♦"
	case Clubs:
		return "♣"
	default:
		panic("invalid card suit")
	}
}
