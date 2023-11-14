package main

import (
	"echo/app"
	"echo/utils"
)

func main() {
	err := utils.Migrate(utils.GetMigrationDbConfig(), "./migrations")
	if err != nil {
		panic(err)
	}

	db, err := utils.SetupAppDb(utils.GetAppDbConfig())
	if err != nil {
		panic(err)
	}

	router := app.AppSetup(db)
	router.Run()
}
