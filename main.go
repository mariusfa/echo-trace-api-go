package main

import (
	"echo/rest"
	"echo/biz"
)

func main() {
	healthController := rest.HealthController{}

	userRepository := &biz.UserRepositoryFake{}
	userService := biz.UserService{UserRepository: userRepository}
	userController := rest.UserController{UserService: userService}
	rest.SetupRouter(&healthController, &userController).Run()
}

