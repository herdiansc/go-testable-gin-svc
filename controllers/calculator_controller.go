package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/herdiansc/go-testable-gin-svc/services"
)

var CalculatorController = calculatorController{}

type calculatorController struct{
	c *gin.Context
}

func (controller calculatorController) Divide(c *gin.Context) {
	result := services.CalculatorService.Divide(c.MustGet("payload").([]byte))
	c.JSON(result.Code, result)
}

func (controller calculatorController) Calculate(c *gin.Context) {
	result := services.CalculatorService.Calculate(c.MustGet("payload").([]byte))
	c.JSON(result.Code, result)
}