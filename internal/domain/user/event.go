package user

import (
	"{{.Module}}/internal/eventbus"
)

type (
	EventUserCreated struct {
		ID    string
		Name  string
		Email string
	}
)

const (
	EventTypeUserCreated = "user.created"
)

func (uc EventUserCreated) Type() eventbus.DomainEventType {
	return EventTypeUserCreated
}
