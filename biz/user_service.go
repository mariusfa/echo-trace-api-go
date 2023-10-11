package biz

import (
	"echo/biz/domain"
)

type UserService struct {
	UserRepository UserRepositoryContract
}

func (us *UserService) Register(userRequest domain.UserRequest) error {
	us.UserRepository.Insert(userRequest.ToUser())
	return nil
}