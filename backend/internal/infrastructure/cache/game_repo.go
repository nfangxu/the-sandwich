package cache

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/the-sandwich/backend/internal/domain/game"
)

const matchKeyPrefix = "match:"

type GameRepository struct {
	client *redis.Client
	ctx    context.Context
}

func NewGameRepository(client *redis.Client) *GameRepository {
	return &GameRepository{
		client: client,
		ctx:    context.Background(),
	}
}

func getMatchKey(matchID string) string {
	return matchKeyPrefix + matchID
}

func (r *GameRepository) Save(state *game.GameState) error {
	data, err := json.Marshal(state)
	if err != nil {
		return err
	}
	err = r.client.Set(r.ctx, getMatchKey(state.MatchID), data, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to save game state: %v", err)
	}
	return nil
}

func (r *GameRepository) Load(matchID string) (*game.GameState, error) {
	data, err := r.client.Get(r.ctx, getMatchKey(matchID)).Bytes()
	if err == redis.Nil {
		return nil, fmt.Errorf("match not found")
	} else if err != nil {
		return nil, fmt.Errorf("failed to load game state: %v", err)
	}

	var state game.GameState
	err = json.Unmarshal(data, &state)
	if err != nil {
		return nil, err
	}
	return &state, nil
}
