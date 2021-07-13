package user

import (
	"github.com/jacexh/golang-ddd-template/internal/eventbus"
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
