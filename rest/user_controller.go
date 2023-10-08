package rest

import (
	"github.com/gin-gonic/gin"
	"echo/biz"
	"echo/rest/dto"
)

type UserController struct {
	UserService *biz.UserService
}

func (uc *UserController) Register(c *gin.Context) {
	var userRequestDTO dto.UserRequestDTO
	if err := c.ShouldBindJSON(&userRequestDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := uc.UserService.Register(userRequestDTO.ToDomain())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
}
