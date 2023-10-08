package rest

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(healthController *HealthController, userController *UserController) *gin.Engine {
	r := gin.Default()
	r.GET("/health", healthController.HealthCheck)

	r.POST("/user/register", userController.Register)
	return r
}
