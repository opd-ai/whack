package procgen

// AnnouncerData holds generated announcer voice data.
type AnnouncerData struct {
	Phrases []AnnouncerPhrase
}

// AnnouncerPhrase represents a generated announcer phrase.
type AnnouncerPhrase struct {
	Text string
	Seed int64
}

// AnnouncerGenerator generates announcer voice data from a seed.
type AnnouncerGenerator struct{}

// Generate produces announcer data deterministically from a seed and genre.
func (g *AnnouncerGenerator) Generate(seed int64, params GenerationParams) (interface{}, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return &AnnouncerData{}, nil
}

// Validate checks that the generator is properly configured.
func (g *AnnouncerGenerator) Validate() error { return nil }
