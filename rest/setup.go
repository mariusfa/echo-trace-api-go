package rest

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(healthController *HealthController) *gin.Engine {
	r := gin.Default()
	r.GET("/health", healthController.HealthCheck)
	return r
}
