package service

import "{{.Module}}/types/repository"

type (
	userService struct {
		repo repository.UserRepository
	}
)
