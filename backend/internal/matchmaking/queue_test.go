package matchmaking

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	redispkg "github.com/redis/go-redis/v9"
	"github.com/the-sandwich/backend/internal/redis"
)

func TestMatchmakingQueue(t *testing.T) {
    // Setup miniredis
    mr, err := miniredis.Run()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer mr.Close()

    redis.Client = redispkg.NewClient(&redispkg.Options{
        Addr: mr.Addr(),
    })

    // Test joining queue
    err = JoinQueue("user1")
    if err != nil {
        t.Fatalf("Expected nil, got %v", err)
    }
    
    JoinQueue("user2")
    
    // Check match creation logic
    players, matched, err := TryCreateMatch()
    if err != nil {
        t.Fatalf("Expected nil, got %v", err)
    }
    if matched {
        t.Errorf("Expected false since only 2 players are in queue")
    }

    JoinQueue("user3")
    
    players, matched, err = TryCreateMatch()
    if err != nil {
        t.Fatalf("Expected nil, got %v", err)
    }
    if !matched {
        t.Errorf("Expected true since 3 players are in queue")
    }
    if len(players) != 3 {
        t.Errorf("Expected 3 players, got %d", len(players))
    }
}
