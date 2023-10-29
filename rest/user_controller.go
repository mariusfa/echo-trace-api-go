package rest

import (
	"echo/biz"
	"echo/rest/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService biz.UserService
}

func (uc *UserController) Register(c *gin.Context) {
	var userRequestDTO dto.UserRequestDTO
	if err := c.ShouldBindJSON(&userRequestDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := uc.UserService.Register(userRequestDTO.ToDomain())

	if err != nil && err.Error() == "User already exists" {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (uc *UserController) Login(c *gin.Context) {
	var userRequestDTO dto.UserRequestDTO
	if err := c.ShouldBindJSON(&userRequestDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := uc.UserService.Login(userRequestDTO.ToDomain())

	if err != nil && err.Error() == "Invalid credentials" {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
}

func (uc *UserController) Validate(c *gin.Context) {
	// TODO: Implement
}
