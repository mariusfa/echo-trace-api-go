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

	userRepository := &biz.UserRepositoryFake{}
	rest.SetupServicesControllers(userRepository).Run()
}
