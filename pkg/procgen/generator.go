// Package procgen provides procedural content generation interfaces and implementations.
package procgen

import "fmt"

// GenreID identifies a game genre.
type GenreID string

const (
	GenreFantasy       GenreID = "fantasy"
	GenreSciFi         GenreID = "sci-fi"
	GenreHorror        GenreID = "horror"
	GenreCyberpunk     GenreID = "cyberpunk"
	GenrePostApocalyptic GenreID = "post-apocalyptic"
)

// ValidGenres contains all valid genre identifiers.
var ValidGenres = []GenreID{
	GenreFantasy, GenreSciFi, GenreHorror, GenreCyberpunk, GenrePostApocalyptic,
}

// GenerationParams holds parameters for procedural generation.
type GenerationParams struct {
	GenreID GenreID
}

// Validate checks that GenerationParams are valid.
func (p GenerationParams) Validate() error {
	for _, g := range ValidGenres {
		if p.GenreID == g {
			return nil
		}
	}
	return &InvalidGenreError{Genre: p.GenreID}
}

// InvalidGenreError is returned when an invalid genre is used.
type InvalidGenreError struct {
	Genre GenreID
}

func (e *InvalidGenreError) Error() string {
	return fmt.Sprintf("invalid genre: %q", e.Genre)
}

// Generator is the PCG interface for all procedural generators.
type Generator interface {
	Generate(seed int64, params GenerationParams) (interface{}, error)
	Validate() error
}
