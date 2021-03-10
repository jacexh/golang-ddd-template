package user

import (
	"context"
)

type (
	Event interface {
		Type() string
	}

	UserCreated struct {
		ID    string
		Name  string
		Email string
	}

	UserDeleted struct {
		ID string
	}

	EventPublisher interface {
		Publish(ctx context.Context, event Event)
		Subscribe(string, Subscriber)
	}

	Subscriber interface {
		Handle(ctx context.Context, event Event) error
	}
)

func (uc UserCreated) Type() string {
	return "user.created"
}

func (ud UserDeleted) Type() string {
	return "user.deleted"
}
