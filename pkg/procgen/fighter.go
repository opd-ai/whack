package procgen

// FighterData holds generated fighter data.
type FighterData struct {
	Name   string
	Speed  float64
	Weight float64
	Jump   float64
	Reach  float64
}

// FighterGenerator generates fighter data from a seed.
type FighterGenerator struct{}

// Generate produces fighter data deterministically from a seed and genre.
func (g *FighterGenerator) Generate(seed int64, params GenerationParams) (interface{}, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return &FighterData{
		Name:   "Fighter",
		Speed:  1.0,
		Weight: 1.0,
		Jump:   1.0,
		Reach:  1.0,
	}, nil
}

// Validate checks that the generator is properly configured.
func (g *FighterGenerator) Validate() error { return nil }
