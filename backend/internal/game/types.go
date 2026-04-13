package game

type Suit string
type Rank int

const (
	Spades   Suit = "Spades"
	Hearts   Suit = "Hearts"
	Clubs    Suit = "Clubs"
	Diamonds Suit = "Diamonds"
	Joker    Suit = "Joker"
)

type Card struct {
	Suit Suit
	Rank Rank // 2-14 (11=J, 12=Q, 13=K, 14=A), 15=SmallJoker, 16=BigJoker
}

type HandType int

const (
	HighCard HandType = iota
	Pair
	Straight
	Flush
	StraightFlush
	Leopard // Three of a kind
)

// Represents the evaluation of 3 cards
type HandResult struct {
	Type  HandType
	Cards []Card // Sorted for comparison
	Score int    // Internal score for easy comparison
}
