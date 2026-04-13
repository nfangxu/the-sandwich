package game

import "testing"

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	if len(deck) != 54 {
		t.Errorf("Expected 54 cards, got %d", len(deck))
	}
}
