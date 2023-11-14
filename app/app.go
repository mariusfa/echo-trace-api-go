package app

import (
	"database/sql"
	"echo/biz/repositories"
	"echo/biz/services"
	"echo/rest/controllers"
	"echo/rest/routes"

	"github.com/gin-gonic/gin"
)

func AppSetup(db *sql.DB) *gin.Engine {
	repositories := repositories.NewRepositories(db)
	services := services.NewServices(repositories)
	controllers := controllers.NewControllers(services)
	routes := routes.SetupRoutes(controllers)
	return routes
}

func AppTestSetup() (*gin.Engine, repositories.Repositories) {
	repositories := repositories.NewRepositoriesFake()
	services := services.NewServices(repositories)
	controllers := controllers.NewControllers(services)
	routes := routes.SetupRoutes(controllers)
	return routes, repositories
}