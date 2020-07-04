package service

import "github.com/jacexh/golang-ddd-template/types/repository"

type (
	userService struct {
		repo repository.UserRepository
	}
)
