// Package combat provides frame-data-driven combat, hitbox resolution, and knockback systems.
package combat

import "github.com/opd-ai/whack/pkg/engine"

// CombatSystem handles hitbox vs hurtbox resolution and damage.
type CombatSystem struct{}

// Update runs combat resolution for one tick.
func (s *CombatSystem) Update(w *engine.World, dt float64) {
	// Skeleton: hitbox vs hurtbox intersection will be implemented here.
}

// KnockbackSystem handles knockback vectors and DI.
type KnockbackSystem struct{}

// Update runs knockback calculations for one tick.
func (s *KnockbackSystem) Update(w *engine.World, dt float64) {
	// Skeleton: knockback vector computation will be implemented here.
}
