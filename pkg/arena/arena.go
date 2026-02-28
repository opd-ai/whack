// Package arena provides platform layout and hazard placement.
package arena

// Arena represents a generated arena.
type Arena struct {
	Platforms []Platform
	Hazards   []Hazard
	BlastZone BlastZone
}

// Platform represents a platform in the arena.
type Platform struct {
	X, Y, W, H float64
	Moving      bool
}

// Hazard represents an arena hazard.
type Hazard struct {
	X, Y   float64
	Type   string
	Active bool
}

// BlastZone defines the arena kill boundaries.
type BlastZone struct {
	Left, Right, Top, Bottom float64
}

// NewArena creates a new arena with a default ground platform.
func NewArena() *Arena {
	return &Arena{
		Platforms: []Platform{
			{X: 200, Y: 400, W: 400, H: 20},
		},
		BlastZone: BlastZone{Left: -100, Right: 900, Top: -100, Bottom: 700},
	}
}
