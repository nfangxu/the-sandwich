package matchmaking

// MatchmakingRepository defines the interface for matchmaking persistence
type MatchmakingRepository interface {
	JoinQueue(userID string) error
	TryCreateMatch() ([]string, bool, error)
}
