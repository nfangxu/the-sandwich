package game

import (
	"math/rand"
	"time"
)

func NewDeck() []Card {
	var deck []Card
	suits := []Suit{Spades, Hearts, Clubs, Diamonds}
	for _, suit := range suits {
		for rank := 2; rank <= 14; rank++ {
			deck = append(deck, Card{Suit: suit, Rank: Rank(rank)})
		}
	}
	deck = append(deck, Card{Suit: Joker, Rank: 15}) // Small Joker
	deck = append(deck, Card{Suit: Joker, Rank: 16}) // Big Joker

	// Shuffle
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})

	return deck
}
