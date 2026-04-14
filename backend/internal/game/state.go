package game

import (
	"encoding/json"
	"fmt"

	redispkg "github.com/redis/go-redis/v9"
	"github.com/the-sandwich/backend/internal/redis"
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

func getMatchKey(matchID string) string {
	return "match:" + matchID
}

// SaveGameState saves the game state to redis
func SaveGameState(state *GameState) error {
	data, err := json.Marshal(state)
	if err != nil {
		return err
	}
	err = redis.Client.Set(redis.Ctx, getMatchKey(state.MatchID), data, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to save game state: %v", err)
	}
	return nil
}

// LoadGameState loads the game state from redis
func LoadGameState(matchID string) (*GameState, error) {
	data, err := redis.Client.Get(redis.Ctx, getMatchKey(matchID)).Bytes()
	if err == redispkg.Nil {
		return nil, fmt.Errorf("match not found")
	} else if err != nil {
		return nil, fmt.Errorf("failed to load game state: %v", err)
	}

	var state GameState
	err = json.Unmarshal(data, &state)
	if err != nil {
		return nil, err
	}
	return &state, nil
}
