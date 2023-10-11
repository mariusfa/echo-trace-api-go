package main

import (
	"echo/biz"
	"echo/rest"
)

func main() {
	healthController := rest.HealthController{}

	userRepository := &biz.UserRepositoryFake{}
	userService := biz.UserService{UserRepository: userRepository}
	userController := rest.UserController{UserService: userService}
	rest.SetupRouter(&healthController, &userController).Run()
}
