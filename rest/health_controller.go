package rest

import (
	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (h *HealthController) HealthCheck(c *gin.Context) {
	c.String(200, "Hello World")
}
