// Package audio provides procedural audio generation including oscillators, envelopes, and motif system.
package audio

// Oscillator generates audio waveforms.
type Oscillator struct {
	Frequency float64
	Amplitude float64
	Waveform  WaveformType
}

// WaveformType identifies a waveform shape.
type WaveformType int

const (
	Sine WaveformType = iota
	Square
	Sawtooth
	Triangle
	Noise
)

// Envelope defines an ADSR amplitude envelope.
type Envelope struct {
	Attack  float64
	Decay   float64
	Sustain float64
	Release float64
}

// Motif represents a musical motif as a sequence of notes.
type Motif struct {
	Notes []Note
	BPM   int
}

// Note represents a single musical note.
type Note struct {
	Pitch    float64
	Duration float64
	Velocity float64
}

// AudioSystem manages audio playback.
type AudioSystem struct {
	// Skeleton: will hold audio context and playback state.
}

// NewAudioSystem creates a new audio system.
func NewAudioSystem() *AudioSystem {
	return &AudioSystem{}
}

// Update runs audio processing for one tick.
func (s *AudioSystem) Update() {
	// Skeleton: audio mixing and playback will be implemented here.
}
