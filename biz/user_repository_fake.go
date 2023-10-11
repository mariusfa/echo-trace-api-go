package biz

import (
	"echo/biz/domain"
)

type UserRepositoryFake struct {
	Users []domain.User
}

func (ur *UserRepositoryFake) Insert(user domain.User) error {
	ur.Users = append(ur.Users, user)
	return nil
}