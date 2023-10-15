package main

import (
	"echo/biz"
	"echo/rest"
	"echo/utils"
)

func main() {
	migrationDbConfig := utils.GetMigrationDbConfig()
	err := utils.Migrate(migrationDbConfig, "./migrations")
	if err != nil {
		panic(err)
	}

	healthController := rest.HealthController{}

	userRepository := &biz.UserRepositoryFake{}
	userService := biz.UserService{UserRepository: userRepository}
	userController := rest.UserController{UserService: userService}
	rest.SetupRouter(&healthController, &userController).Run()
}
