package user

import (
	"errors"

	"github.com/jacexh/golang-ddd-template/internal/domain/event"
)

type UserEntity struct {
	ID       string
	Name     string
	Password string
	Email    string
	Events   *event.Events // https://docs.microsoft.com/en-us/dotnet/architecture/microservices/microservice-ddd-cqrs-patterns/domain-events-design-implementation
}

// todo
func genUserID() string {
	return ""
}

func NewUser(name, password, email string) *UserEntity {
	u := &UserEntity{
		ID:       genUserID(),
		Name:     name,
		Password: password,
		Email:    email,
		Events:   event.NewEvents(),
	}
	u.Events.Add(EventUserCreated{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	})
	return u
}

// Validate 参数校验
func (u *UserEntity) Validate() error {
	if u.Name == "" || u.Email == "" {
		return errors.New("bad user information")
	}
	return nil
}
