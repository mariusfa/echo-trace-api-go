package biz

import (
	"echo/biz/domain"
	"errors"
)

type UserRepositoryFake struct {
	Users []domain.User
}

func (ur *UserRepositoryFake) Insert(user domain.User) error {
	ur.Users = append(ur.Users, user)
	return nil
}

func (ur *UserRepositoryFake) GetByName(name string) (domain.User, error) {
	for _, user := range ur.Users {
		if user.Name == name {
			return user, nil
		}
	}
	return domain.User{}, errors.New("User not found")
}
