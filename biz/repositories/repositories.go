package repositories

import (
	"database/sql"
	"echo/biz"
)

type Repositories struct {
	UserRepository biz.UserRepositoryContract
}

func NewRepositories(db *sql.DB) Repositories {
	return Repositories{
		UserRepository: biz.NewUserRepository(db),
	}
}

func NewRepositoriesFake() Repositories {
	return Repositories{
		UserRepository: biz.NewUserRepositoryFake(),
	}
}

