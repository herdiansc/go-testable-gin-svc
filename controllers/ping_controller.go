package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/herdiansc/go-testable-gin-svc/services"
)

var PingController = pingController{}

type pingController struct{
	c *gin.Context
}

func (controller pingController) Pong(c *gin.Context) {
	result, err := services.PingService.Pong()
	code := 200
	message := result
	if err != nil {
		code = 500
		message = err.Error()
	}

	c.JSON(code, gin.H{
		"message": message,
	})
}

func (controller pingController) Divide(c *gin.Context) {
	result := services.PingService.Divide(c.MustGet("payload").([]byte))
	c.JSON(result.Code, result)
}

func (controller pingController) Calculate(c *gin.Context) {
	result := services.PingService.Calculate(c.MustGet("payload").([]byte))
	c.JSON(result.Code, result)
}