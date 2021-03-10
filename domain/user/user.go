package user

type UserEntity struct {
	ID       string
	Name     string
	Password string
	Email    string
	Changes  []Event // https://docs.microsoft.com/en-us/dotnet/architecture/microservices/microservice-ddd-cqrs-patterns/domain-events-design-implementation
}

func NewUser(id, name, password, email string) *UserEntity {
	u := &UserEntity{
		ID:       id,
		Name:     name,
		Password: password,
		Email:    email,
		Changes:  make([]Event, 0),
	}
	u.Changes = append(u.Changes, UserCreated{
		ID:    id,
		Name:  u.Name,
		Email: u.Email,
	})
	return u
}
