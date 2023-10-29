package rest

import (
	"echo/biz"

	"github.com/gin-gonic/gin"
)

func SetupServices(userRepo biz.UserRepositoryContract) biz.UserService {
	// TODO impl
}

func SetupServicesControllers(userRepo biz.UserRepositoryContract) *gin.Engine {

	userService := biz.UserService{UserRepository: userRepo}
	userController := UserController{UserService: userService}
	healthController := HealthController{}
	return SetupRouter(healthController, userController)
}

func SetupRouter(healthController HealthController, userController UserController) *gin.Engine {
	r := gin.Default()
	r.GET("/health", healthController.HealthCheck)

	r.POST("/user/register", userController.Register)
	return r
}
