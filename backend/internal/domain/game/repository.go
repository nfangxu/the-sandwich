package game

// GameRepository defines the interface for game state persistence
type GameRepository interface {
	Save(state *GameState) error
	Load(matchID string) (*GameState, error)
}
