package procgen

// ArenaData holds generated arena layout data.
type ArenaData struct {
	Platforms []Platform
	BlastZone BlastZone
}

// Platform represents a platform in the arena.
type Platform struct {
	X, Y, W, H float64
}

// BlastZone defines the arena boundaries.
type BlastZone struct {
	Left, Right, Top, Bottom float64
}

// ArenaGenerator generates arena layouts from a seed.
type ArenaGenerator struct{}

// Generate produces arena data deterministically from a seed and genre.
func (g *ArenaGenerator) Generate(seed int64, params GenerationParams) (interface{}, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return &ArenaData{
		Platforms: []Platform{
			{X: 200, Y: 400, W: 400, H: 20},
		},
		BlastZone: BlastZone{Left: -100, Right: 900, Top: -100, Bottom: 700},
	}, nil
}

// Validate checks that the generator is properly configured.
func (g *ArenaGenerator) Validate() error { return nil }
