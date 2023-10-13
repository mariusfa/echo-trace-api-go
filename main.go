package main

import (
	"echo/biz"
	"echo/rest"
	"echo/utils"
)

func main() {
	if (utils.GetEnv("APP_ENV", "prod") == "dev") {
		utils.MigrateBase()
	}

	healthController := rest.HealthController{}

	userRepository := &biz.UserRepositoryFake{}
	userService := biz.UserService{UserRepository: userRepository}
	userController := rest.UserController{UserService: userService}
	rest.SetupRouter(&healthController, &userController).Run()
}
