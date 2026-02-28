package procgen

// MovesetData holds generated moveset data.
type MovesetData struct {
	Normals  []Move
	Specials []Move
	Supers   []Move
}

// Move represents a single move with frame data.
type Move struct {
	Name     string
	Startup  int
	Active   int
	Recovery int
	Damage   float64
}

// MovesetGenerator generates movesets from a seed.
type MovesetGenerator struct{}

// Generate produces moveset data deterministically from a seed and genre.
func (g *MovesetGenerator) Generate(seed int64, params GenerationParams) (interface{}, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return &MovesetData{}, nil
}

// Validate checks that the generator is properly configured.
func (g *MovesetGenerator) Validate() error { return nil }
