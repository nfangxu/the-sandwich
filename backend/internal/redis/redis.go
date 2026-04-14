package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client
var Ctx = context.Background()

// ConnectRedis initializes the Redis connection
func ConnectRedis(addr string, password string, db int) error {
	Client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	_, err := Client.Ping(Ctx).Result()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %v", err)
	}
	
	log.Println("Redis connection established")
	return nil
}
