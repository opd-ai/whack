package engine

// Scheduler manages system execution order.
type Scheduler struct {
	systems []System
}

// NewScheduler creates a new scheduler.
func NewScheduler() *Scheduler {
	return &Scheduler{}
}

// Register adds a system to the scheduler.
func (s *Scheduler) Register(sys System) {
	s.systems = append(s.systems, sys)
}

// Update runs all registered systems.
func (s *Scheduler) Update(w *World, dt float64) {
	for _, sys := range s.systems {
		sys.Update(w, dt)
	}
}
