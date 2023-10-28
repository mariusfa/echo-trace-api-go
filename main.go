package main

import (
	"echo/rest"
	"echo/utils"
	"net/http"
)

func main() {
	migrationDbConfig := utils.GetMigrationDbConfig()
	err := utils.Migrate(migrationDbConfig, "./migrations")
	if err != nil {
		panic(err)
	}

	healthController := rest.HealthController{}
	server := rest.SetupHandlers(healthController)
	http.ListenAndServe(":8080", server)

	// userRepository := &biz.UserRepositoryFake{}
	// rest.SetupServicesControllers(userRepository).Run()
}
