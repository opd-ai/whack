// Package rendering provides runtime sprite generation, palette management, and post-processing.
package rendering

import "image/color"

// Palette represents a genre-specific color palette.
type Palette struct {
	Primary   color.RGBA
	Secondary color.RGBA
	Accent    color.RGBA
	BG        color.RGBA
}

// SpriteRenderer generates and renders sprites at runtime.
type SpriteRenderer struct {
	// Skeleton: will hold sprite generation state.
}

// NewSpriteRenderer creates a new sprite renderer.
func NewSpriteRenderer() *SpriteRenderer {
	return &SpriteRenderer{}
}

// PostProcessor applies visual post-processing effects.
type PostProcessor struct {
	ScreenShake         bool
	Scanlines           bool
	ChromaticAberration bool
	BloodTint           bool
	Glow                bool
}

// NewPostProcessor creates a new post-processor.
func NewPostProcessor() *PostProcessor {
	return &PostProcessor{}
}
