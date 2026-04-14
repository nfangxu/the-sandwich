package game

import (
	"time"
)

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

// PlayCards handles moving cards from hand to played cards for a round.
func PlayCards(state *GameState, userID string, cardIndices []int) error {
	for i := range state.Players {
		if state.Players[i].UserID == userID {
			if state.Players[i].HasPlayed {
				return nil // already played
			}
			
			var played []Card
			var newHand []Card
			indexMap := make(map[int]bool)
			for _, idx := range cardIndices {
				indexMap[idx] = true
			}
			
			for idx, card := range state.Players[i].Hand {
				if indexMap[idx] {
					played = append(played, card)
				} else {
					newHand = append(newHand, card)
				}
			}
			
			state.Players[i].Hand = newHand
			state.Players[i].PlayedCards = played
			state.Players[i].HasPlayed = true
			return nil
		}
	}
	return nil
}
