package user

import "github.com/jacexh/golang-ddd-template/internal/domain/event"

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

func (uc EventUserCreated) Type() event.DomainEventType {
	return EventTypeUserCreated
}
