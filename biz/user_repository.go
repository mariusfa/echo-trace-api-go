package biz

import (
	"echo/biz/domain"
)

type User struct {
	Id             int64
	Name           string
	HashedPassword string
	ApiToken       string
}

type UserRepositoryContract interface {
	Insert(user domain.User) error
	GetByName(name string) (domain.User, error)
}
