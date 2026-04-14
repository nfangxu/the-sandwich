package game

import (
	"github.com/the-sandwich/backend/internal/domain/game"
)

// GameService is the application service for game operations
type GameService struct {
	repo game.GameRepository
}

// NewGameService creates a new game application service
func NewGameService(repo game.GameRepository) *GameService {
	return &GameService{repo: repo}
}

// CreateMatch creates a new match with the given players
func (s *GameService) CreateMatch(matchID string, playerIDs []string) (*game.GameState, error) {
	state := game.InitMatch(matchID, playerIDs)
	if err := s.repo.Save(state); err != nil {
		return nil, err
	}
	return state, nil
}

// GetMatch retrieves a match by ID
func (s *GameService) GetMatch(matchID string) (*game.GameState, error) {
	return s.repo.Load(matchID)
}

// PlayCards processes a player's card play
func (s *GameService) PlayCards(matchID, userID string, cardIndices []int) (*game.GameState, error) {
	state, err := s.repo.Load(matchID)
	if err != nil {
		return nil, err
	}

	if err := game.PlayCards(state, userID, cardIndices); err != nil {
		return nil, err
	}

	if err := s.repo.Save(state); err != nil {
		return nil, err
	}

	return state, nil
}

// AdvanceRound advances the game to the next round
func (s *GameService) AdvanceRound(matchID string) (*game.GameState, error) {
	state, err := s.repo.Load(matchID)
	if err != nil {
		return nil, err
	}

	game.AdvanceRound(state)

	if err := s.repo.Save(state); err != nil {
		return nil, err
	}

	return state, nil
}
