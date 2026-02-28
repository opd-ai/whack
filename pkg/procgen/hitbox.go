package procgen

// HitboxData holds generated hitbox shapes.
type HitboxData struct {
	Shapes []HitboxShape
}

// HitboxShape represents a hitbox shape for a move.
type HitboxShape struct {
	X, Y, W, H  float64
	ActiveFrames int
}

// HitboxGenerator generates hitbox data from a seed.
type HitboxGenerator struct{}

// Generate produces hitbox data deterministically from a seed and genre.
func (g *HitboxGenerator) Generate(seed int64, params GenerationParams) (interface{}, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return &HitboxData{}, nil
}

// Validate checks that the generator is properly configured.
func (g *HitboxGenerator) Validate() error { return nil }
