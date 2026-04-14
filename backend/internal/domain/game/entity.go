package game

import (
	"time"
)

type PlayerState struct {
	UserID       string
	Hand         []Card
	PlayedCards  []Card
	Score        int
	HasPlayed    bool
	IsAutoPlayed bool
}

type GameState struct {
	MatchID     string
	Players     []PlayerState
	Round       int // 1 to 5
	Status      string // "WAITING", "PLAYING", "ROUND_OVER", "FINISHED"
	PublicCards []Card // length 4
	Deck        []Card
	TurnExpires int64 // Unix timestamp
}

func InitMatch(matchID string, userIDs []string) *GameState {
	deck := NewDeck()

	// Draw 4 public cards
	publicCards := deck[:4]
	deck = deck[4:]

	players := make([]PlayerState, len(userIDs))
	for i, uid := range userIDs {
		// Draw 5 cards per player
		hand := deck[:5]
		deck = deck[5:]

		players[i] = PlayerState{
			UserID: uid,
			Hand:   hand,
			Score:  0,
		}
	}

	return &GameState{
		MatchID:     matchID,
		Players:     players,
		Round:       1,
		Status:      "PLAYING",
		PublicCards: publicCards,
		Deck:        deck,
		TurnExpires: time.Now().Add(30 * time.Second).Unix(),
	}
}
