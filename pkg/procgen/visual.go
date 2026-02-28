package procgen

// ArenaVisualData holds generated arena visual data.
type ArenaVisualData struct {
	TilePalette []Color
	Background  []BackgroundLayer
}

// FighterVisualData holds generated fighter visual data.
type FighterVisualData struct {
	Palette []Color
	Layers  []SpriteLayer
}

// Color represents an RGBA color.
type Color struct {
	R, G, B, A uint8
}

// BackgroundLayer represents a parallax background layer.
type BackgroundLayer struct {
	Depth float64
	Seed  int64
}

// SpriteLayer represents a visual layer for fighter rendering.
type SpriteLayer struct {
	Name string
	Seed int64
}

// ArenaVisualGenerator generates arena visual data from a seed.
type ArenaVisualGenerator struct{}

// Generate produces arena visual data deterministically from a seed and genre.
func (g *ArenaVisualGenerator) Generate(seed int64, params GenerationParams) (interface{}, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return &ArenaVisualData{}, nil
}

// Validate checks that the generator is properly configured.
func (g *ArenaVisualGenerator) Validate() error { return nil }

// FighterVisualGenerator generates fighter visual data from a seed.
type FighterVisualGenerator struct{}

// Generate produces fighter visual data deterministically from a seed and genre.
func (g *FighterVisualGenerator) Generate(seed int64, params GenerationParams) (interface{}, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	return &FighterVisualData{}, nil
}

// Validate checks that the generator is properly configured.
func (g *FighterVisualGenerator) Validate() error { return nil }
