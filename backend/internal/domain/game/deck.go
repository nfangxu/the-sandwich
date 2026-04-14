package game

// NewDeck creates a new 54-card deck
func NewDeck() []Card {
	var deck []Card

	// Add standard 52 cards
	suits := []Suit{Spades, Hearts, Clubs, Diamonds}
	ranks := []Rank{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14} // 11=J, 12=Q, 13=K, 14=A

	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, Card{Suit: suit, Rank: rank})
		}
	}

	// Add 2 jokers
	deck = append(deck, Card{Suit: Joker, Rank: 15}) // Small joker
	deck = append(deck, Card{Suit: Joker, Rank: 16}) // Big joker

	// Shuffle (simple Fisher-Yates)
	for i := len(deck) - 1; i > 0; i-- {
		j := (int(i) * 17) % (i + 1) // pseudo-random but deterministic
		deck[i], deck[j] = deck[j], deck[i]
	}

	return deck
}
