package game

import (
	"sort"
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

func CompareHands(h1, h2 HandResult) int {
	if h1.Type > h2.Type {
		return 1
	} else if h1.Type < h2.Type {
		return -1
	}
	// For simplicity, we skip rank-by-rank tie-breaking in this prototype, 
    // but in a real app, we'd compare card ranks.
	return 0
}

func AdvanceRound(state *GameState) {
    if state.Status == "FINISHED" {
        return
    }

    // 1. Collect and evaluate all hands
    results := make([]struct {
        UserID string
        Hand   HandResult
    }, len(state.Players))

    for i, p := range state.Players {
        cards := append([]Card{}, p.PlayedCards...)
        // For rounds 1-4, add the public card (P1 to P4)
        if state.Round <= 4 {
            cards = append(cards, state.PublicCards[state.Round-1])
        }
        results[i].UserID = p.UserID
        results[i].Hand = EvaluateHand(cards)
    }

    // 2. Sort by hand strength descending
    sort.Slice(results, func(i, j int) bool {
        return CompareHands(results[i].Hand, results[j].Hand) > 0
    })

    // 3. Identify the Sandwich (2nd place)
    // Rule: "第二名输，赔付给第一名和第三名"
    points := state.Round // Round 1: 1pt, Round 2: 2pt...

    for i := range state.Players {
        if state.Players[i].UserID == results[0].UserID {
            state.Players[i].Score += points
        } else if state.Players[i].UserID == results[1].UserID {
            state.Players[i].Score -= (points * 2)
        } else if state.Players[i].UserID == results[2].UserID {
            state.Players[i].Score += points
        }
    }

    // 4. Reset round state
    for i := range state.Players {
        state.Players[i].HasPlayed = false
        state.Players[i].PlayedCards = nil
    }

    // 5. Refill cards (only rounds 1-3 need refilling, as per rules)
    if state.Round <= 3 {
        for i := range state.Players {
            refillCount := 2
            if len(state.Deck) >= refillCount {
                state.Players[i].Hand = append(state.Players[i].Hand, state.Deck[:refillCount]...)
                state.Deck = state.Deck[refillCount:]
            }
        }
    }

    // 6. Advance round
    if state.Round < 5 {
        state.Round++
        state.TurnExpires = time.Now().Add(30 * time.Second).Unix()
    } else {
        state.Status = "FINISHED"
    }
}
