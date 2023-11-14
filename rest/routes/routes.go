package routes

import (
	"echo/rest/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(controllers controllers.Controllers) *gin.Engine {
	r := gin.Default()
	r.GET("/health", controllers.HealthController.HealthCheck)

	r.POST("/user/register", controllers.UserController.Register)
	return r
}