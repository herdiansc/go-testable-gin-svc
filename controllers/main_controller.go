package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/herdiansc/go-testable-gin-svc/services"
)

var MainController = mainController{}

type mainController struct{
	c *gin.Context
}

func (controller mainController) Pong(c *gin.Context) {
	result := services.MainService.Pong()
	c.JSON(result.Code, result)
}

func (controller mainController) HealthCheck(c *gin.Context) {
	result := services.MainService.HealthCheck()
	c.JSON(result.Code, result)
}