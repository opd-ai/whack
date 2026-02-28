package engine

// Event represents a game event.
type Event struct {
	Type string
	Data interface{}
}

// EventBus provides publish/subscribe event handling.
type EventBus struct {
	handlers map[string][]func(Event)
}

// NewEventBus creates a new event bus.
func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[string][]func(Event)),
	}
}

// Subscribe registers a handler for an event type.
func (b *EventBus) Subscribe(eventType string, handler func(Event)) {
	b.handlers[eventType] = append(b.handlers[eventType], handler)
}

// Publish sends an event to all subscribed handlers.
func (b *EventBus) Publish(e Event) {
	for _, h := range b.handlers[e.Type] {
		h(e)
	}
}
