// Package engine provides the ECS core, event bus, and scheduler.
package engine

// Entity is a unique identifier for an ECS entity.
type Entity uint64

// World manages entities and component storage.
type World struct {
	nextID     Entity
	Positions  *ComponentStore[PositionComponent]
	Hitboxes   *ComponentStore[HitboxComponent]
	Hurtboxes  *ComponentStore[HurtboxComponent]
	Stats      *ComponentStore[FighterStatsComponent]
	Knockbacks *ComponentStore[KnockbackComponent]
}

// NewWorld creates a new ECS world with initialized component stores.
func NewWorld() *World {
	return &World{
		nextID:     1,
		Positions:  NewComponentStore[PositionComponent](),
		Hitboxes:   NewComponentStore[HitboxComponent](),
		Hurtboxes:  NewComponentStore[HurtboxComponent](),
		Stats:      NewComponentStore[FighterStatsComponent](),
		Knockbacks: NewComponentStore[KnockbackComponent](),
	}
}

// NewEntity creates a new entity and returns its ID.
func (w *World) NewEntity() Entity {
	id := w.nextID
	w.nextID++
	return id
}

// ComponentStore is a sparse-set backed dense component array.
type ComponentStore[T any] struct {
	dense    []T
	entities []Entity
	sparse   map[Entity]int
}

// NewComponentStore creates a new component store.
func NewComponentStore[T any]() *ComponentStore[T] {
	return &ComponentStore[T]{
		sparse: make(map[Entity]int),
	}
}

// Set assigns a component value to an entity.
func (s *ComponentStore[T]) Set(e Entity, c T) {
	if idx, ok := s.sparse[e]; ok {
		s.dense[idx] = c
		return
	}
	s.sparse[e] = len(s.dense)
	s.dense = append(s.dense, c)
	s.entities = append(s.entities, e)
}

// Get retrieves the component for an entity.
func (s *ComponentStore[T]) Get(e Entity) (T, bool) {
	idx, ok := s.sparse[e]
	if !ok {
		var zero T
		return zero, false
	}
	return s.dense[idx], true
}

// All returns all dense component data for iteration.
func (s *ComponentStore[T]) All() []T {
	return s.dense
}

// Entities returns the entity list corresponding to dense storage.
func (s *ComponentStore[T]) Entities() []Entity {
	return s.entities
}

// PositionComponent stores entity position.
type PositionComponent struct {
	X, Y float64
}

// HitboxComponent stores active attack hitbox data.
type HitboxComponent struct {
	X, Y, W, H float64
	Damage      float64
	Active      bool
}

// HurtboxComponent stores vulnerable area data.
type HurtboxComponent struct {
	X, Y, W, H float64
}

// FighterStatsComponent stores fighter stat data.
type FighterStatsComponent struct {
	Speed  float64
	Weight float64
	Jump   float64
	Reach  float64
	Damage float64
}

// KnockbackComponent stores knockback state.
type KnockbackComponent struct {
	VX, VY  float64
	Percent float64
}

// System is the interface for ECS systems.
type System interface {
	Update(w *World, dt float64)
}
