package eventbus

import "context"

type (
	Events struct {
		events []DomainEvent
		index  int
	}
)

func NewEvents() *Events {
	return &Events{
		events: make([]DomainEvent, 0),
		index:  0,
	}
}

func (el *Events) Add(event DomainEvent) {
	el.events = append(el.events, event)
}

func (el *Events) next() (DomainEvent, bool) {
	if el.index >= len(el.events) {
		return nil, false
	}
	el.index++
	return el.events[el.index-1], true
}

func (el *Events) Dispatch(ctx context.Context) {
	for ev, got := el.next(); got; {
		Publish(ctx, ev)
	}
}
