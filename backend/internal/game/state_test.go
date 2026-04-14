package game

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	redispkg "github.com/redis/go-redis/v9"
	"github.com/the-sandwich/backend/internal/redis"
)

func TestGameStateSaveLoad(t *testing.T) {
	mr, _ := miniredis.Run()
	defer mr.Close()

	redis.Client = redispkg.NewClient(&redispkg.Options{
		Addr: mr.Addr(),
	})

	state := &GameState{
		MatchID: "match_123",
		Players: []string{"p1", "p2", "p3"},
		Round:   1,
        Status:  "PLAYING",
	}

	err := SaveGameState(state)
	if err != nil {
		t.Fatalf("Failed to save state: %v", err)
	}

	loadedState, err := LoadGameState("match_123")
	if err != nil {
		t.Fatalf("Failed to load state: %v", err)
	}

	if loadedState.MatchID != "match_123" {
		t.Errorf("Expected match_123, got %s", loadedState.MatchID)
	}
	if loadedState.Round != 1 {
		t.Errorf("Expected round 1, got %d", loadedState.Round)
	}
}
