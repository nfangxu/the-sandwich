package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

const QueueKey = "matchmaking:queue"

type MatchmakingRepository struct {
	client *redis.Client
	ctx    context.Context
}

func NewMatchmakingRepository(client *redis.Client) *MatchmakingRepository {
	return &MatchmakingRepository{
		client: client,
		ctx:    context.Background(),
	}
}

func (r *MatchmakingRepository) JoinQueue(userID string) error {
	err := r.client.RPush(r.ctx, QueueKey, userID).Err()
	if err != nil {
		return fmt.Errorf("failed to join queue: %v", err)
	}
	return nil
}

func (r *MatchmakingRepository) TryCreateMatch() ([]string, bool, error) {
	length, err := r.client.LLen(r.ctx, QueueKey).Result()
	if err != nil {
		return nil, false, err
	}

	if length >= 3 {
		var players []string
		for i := 0; i < 3; i++ {
			val, err := r.client.LPop(r.ctx, QueueKey).Result()
			if err != nil {
				return nil, false, err
			}
			players = append(players, val)
		}
		return players, true, nil
	}

	return nil, false, nil
}
