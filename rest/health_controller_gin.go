package rest

import (
	"github.com/gin-gonic/gin"
)

type HealthControllerGin struct{}

func (h *HealthControllerGin) HealthCheck(c *gin.Context) {
	c.String(200, "Hello World")
}
