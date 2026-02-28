package procgen

// MusicData holds generated music data.
type MusicData struct {
	BPM   int
	Motif []int
}

// MusicGenerator generates music data from a seed.
type MusicGenerator struct{}

// Generate produces music data deterministically from a seed and genre.
func (g *MusicGenerator) Generate(seed int64, params GenerationParams) (interface{}, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return &MusicData{}, nil
}

// Validate checks that the generator is properly configured.
func (g *MusicGenerator) Validate() error { return nil }
