package biz

import (
	"echo/biz/domain"
)

type UserService struct {}

func (us *UserService) Register(userRequest domain.UserRequest) error {
	return nil
}