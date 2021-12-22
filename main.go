package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/herdiansc/go-testable-gin-svc/controllers"
	"github.com/herdiansc/go-testable-gin-svc/services"
)

func init() {
	services.MainService.RegisterAppStartTime()
}

func ProcessPayload() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" {
			payload, _ := c.GetRawData()
			c.Set("payload", payload)
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(ProcessPayload())

	r.GET("/main/healthcheck", controllers.MainController.HealthCheck)
	r.GET("/main/ping", controllers.MainController.Pong)
	r.POST("/calculators/divide", controllers.CalculatorController.Divide)
	r.POST("/calculators/calculate", controllers.CalculatorController.Calculate)
	r.Run()

	fmt.Printf("Application Runing...\n")
}
