package controllers

import (
	"echo/biz/services"
	"echo/rest/health"
	"echo/rest/user"
)

type Controllers struct {
	UserController user.UserController
	HealthController health.HealthController
}

func NewControllers(services services.Services) Controllers {
	return Controllers{
		UserController: user.NewUserController(services.UserService),
		HealthController: health.NewHealthController() ,
	}
}