package services

import (
	"echo/biz"
	"echo/biz/repositories"
)

type Services struct {
	UserService biz.UserService
}

func NewServices(repositories repositories.Repositories) Services {
	return Services{
		UserService: biz.NewUserService(repositories.UserRepository),
	}
}
