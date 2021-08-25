package user

import (
	"errors"

	"{{.Module}}/internal/eventbus"
	"golang.org/x/crypto/bcrypt"
)

type UserEntity struct {
	ID             string
	Name           string
	HashedPassword []byte
	Email          string
	Events         *eventbus.Events // https://docs.microsoft.com/en-us/dotnet/architecture/microservices/microservice-ddd-cqrs-patterns/domain-events-design-implementation
}

// todo
func genUserID() string {
	return ""
}

func NewUser(name, password, email string) (*UserEntity, error) {
	u := &UserEntity{
		ID:     genUserID(),
		Name:   name,
		Email:  email,
		Events: eventbus.NewEvents(),
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.HashedPassword = hashed
	u.Events.Add(EventUserCreated{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	})
	return u, nil
}

// Validate 参数校验
func (u *UserEntity) Validate() error {
	if u.Name == "" || u.Email == "" {
		return errors.New("bad user information")
	}
	return nil
}

func (u *UserEntity) CheckPassword(pwd string) error {
	return bcrypt.CompareHashAndPassword(u.HashedPassword, []byte(pwd))
}
