package main

import (
	"echo/rest"
	"echo/biz"
)

func main() {
	healthController := &rest.HealthController{}

	userService := &biz.UserService{}
	userController := &rest.UserController{UserService: userService}
	rest.SetupRouter(healthController, userController).Run()
}

