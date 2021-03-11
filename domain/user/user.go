package user

import "github.com/jacexh/golang-ddd-template/domain/event"

type UserEntity struct {
	ID       string
	Name     string
	Password string
	Email    string
	Events   *event.Events // https://docs.microsoft.com/en-us/dotnet/architecture/microservices/microservice-ddd-cqrs-patterns/domain-events-design-implementation
}

func NewUser(id, name, password, email string) *UserEntity {
	u := &UserEntity{
		ID:       id,
		Name:     name,
		Password: password,
		Email:    email,
		Events:   event.NewEvents(),
	}
	u.Events.Add(EventUserCreated{
		ID:    id,
		Name:  u.Name,
		Email: u.Email,
	})
	return u
}
