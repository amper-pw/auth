package controller

import (
	"github.com/amper-pw/auth/pkg/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	services *service.Service
}

func NewControllers(services *service.Service) *Controller {
	return &Controller{
		services: services,
	}
}

func (c *Controller) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", c.signUp)
		auth.POST("/sign-in", c.signIn)
	}

	return router
}
