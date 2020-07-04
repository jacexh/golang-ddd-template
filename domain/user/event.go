package user

type UserEvent int

const (
	UserCreated UserEvent = iota + 1
	UserDeleted
)
