// Package physics provides platform collision, gravity, and movement systems.
package physics

import "github.com/opd-ai/whack/pkg/engine"

// PhysicsSystem handles movement, gravity, and platform collision.
type PhysicsSystem struct {
	Gravity float64
}

// NewPhysicsSystem creates a new physics system with default gravity.
func NewPhysicsSystem() *PhysicsSystem {
	return &PhysicsSystem{Gravity: 0.5}
}

// Update runs physics simulation for one tick.
func (s *PhysicsSystem) Update(w *engine.World, dt float64) {
	// Skeleton: gravity, collision, and movement will be implemented here.
}
