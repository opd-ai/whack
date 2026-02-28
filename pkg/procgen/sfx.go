package procgen

// SFXData holds generated sound effect data.
type SFXData struct {
	Samples []SFXSample
}

// SFXSample represents a generated sound effect sample.
type SFXSample struct {
	Name string
	Seed int64
}

// SFXGenerator generates sound effect data from a seed.
type SFXGenerator struct{}

// Generate produces SFX data deterministically from a seed and genre.
func (g *SFXGenerator) Generate(seed int64, params GenerationParams) (interface{}, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return &SFXData{}, nil
}

// Validate checks that the generator is properly configured.
func (g *SFXGenerator) Validate() error { return nil }
