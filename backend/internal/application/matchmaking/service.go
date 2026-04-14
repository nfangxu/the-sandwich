package matchmaking

import (
	"github.com/the-sandwich/backend/internal/domain/matchmaking"
)

// MatchmakingService is the application service for matchmaking operations
type MatchmakingService struct {
	repo matchmaking.MatchmakingRepository
}

// NewMatchmakingService creates a new matchmaking application service
func NewMatchmakingService(repo matchmaking.MatchmakingRepository) *MatchmakingService {
	return &MatchmakingService{repo: repo}
}

// JoinQueue adds a user to the matchmaking queue
func (s *MatchmakingService) JoinQueue(userID string) error {
	return s.repo.JoinQueue(userID)
}

// TryCreateMatch attempts to create a match from queued players
func (s *MatchmakingService) TryCreateMatch() ([]string, bool, error) {
	return s.repo.TryCreateMatch()
}
