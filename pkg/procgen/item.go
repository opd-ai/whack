package procgen

// ItemData holds generated item data.
type ItemData struct {
	Name     string
	Type     string
	StatMods map[string]float64
}

// ItemGenerator generates items from a seed.
type ItemGenerator struct{}

// Generate produces item data deterministically from a seed and genre.
func (g *ItemGenerator) Generate(seed int64, params GenerationParams) (interface{}, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return &ItemData{
		StatMods: make(map[string]float64),
	}, nil
}

// Validate checks that the generator is properly configured.
func (g *ItemGenerator) Validate() error { return nil }
