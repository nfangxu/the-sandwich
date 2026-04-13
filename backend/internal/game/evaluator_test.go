package game

import "testing"

func TestEvaluateHand(t *testing.T) {
	cards := []Card{
		{Suit: Spades, Rank: 14},
		{Suit: Hearts, Rank: 14},
		{Suit: Joker, Rank: 15}, // Small Joker acts as 14 (A)
	}
	result := EvaluateHand(cards)
	if result.Type != Leopard {
		t.Errorf("Expected Leopard, got %v", result.Type)
	}
}
