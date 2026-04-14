package matchmaking

import (
	"fmt"

	redispkg "github.com/redis/go-redis/v9"
	"github.com/the-sandwich/backend/internal/redis"
)

const QueueKey = "matchmaking_queue"

// JoinQueue adds a user to the matchmaking queue
func JoinQueue(userID string) error {
	// We use ZADD or SADD to prevent duplicates, but a simple List is fine if we check existence.
    // For simplicity, we just RPUSH. In real app, ensure they aren't already in queue.
	err := redis.Client.RPush(redis.Ctx, QueueKey, userID).Err()
	if err != nil {
		return fmt.Errorf("failed to join queue: %v", err)
	}
	return nil
}

// TryCreateMatch checks if there are enough players (3) and removes them to form a match
func TryCreateMatch() ([]string, bool, error) {
    // We need 3 players. We can use LPOP with count in Redis >= 6.2, or just LLEN + LPOP.
    length, err := redis.Client.LLen(redis.Ctx, QueueKey).Result()
    if err != nil && err != redispkg.Nil {
        return nil, false, err
    }

    if length >= 3 {
        // Pop 3 players
        // Using a transaction to ensure atomic pop of exactly 3 if needed,
        // but sequential LPOP is fine for this minimal prototype.
        var players []string
        for i := 0; i < 3; i++ {
            val, err := redis.Client.LPop(redis.Ctx, QueueKey).Result()
            if err != nil {
                // If this happens, we might have lost a player mid-pop.
                // A Lua script is safer, but we keep it simple here.
                return nil, false, err
            }
            players = append(players, val)
        }
        return players, true, nil
    }

    return nil, false, nil
}
