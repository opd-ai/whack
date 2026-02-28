// Package tournament provides bracket generation and spectator state.
package tournament

// Bracket represents a tournament bracket.
type Bracket struct {
	Type    BracketType
	Players []string
	Rounds  []Round
}

// BracketType identifies the bracket format.
type BracketType string

const (
	SingleElimination BracketType = "single"
	DoubleElimination BracketType = "double"
)

// Round represents a round in the tournament.
type Round struct {
	Matches []Match
}

// Match represents a single tournament match.
type Match struct {
	Player1 string
	Player2 string
	Winner  string
}

// TournamentGenerator generates tournament brackets deterministically from a seed.
type TournamentGenerator struct{}

// Generate produces a bracket for the given players using the provided seed for deterministic ordering.
func (g *TournamentGenerator) Generate(seed int64, players []string, bracketType BracketType) *Bracket {
	_ = seed // Skeleton: seed will drive deterministic bracket ordering.
	return &Bracket{
		Type:    bracketType,
		Players: players,
	}
}

// SpectatorState holds read-only state for spectators.
type SpectatorState struct {
	Frame    uint32
	Entities []SpectatorEntity
}

// SpectatorEntity holds entity state for spectator rendering.
type SpectatorEntity struct {
	ID   uint64
	X, Y float64
}
