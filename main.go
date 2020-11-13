package main

import (
	"github.com/gin-gonic/gin"
	"github.com/herdiansc/go-testable-gin-svc/controllers"
)

func ProcessPayload() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "POST" {
			payload, _ := c.GetRawData()
			//fmt.Printf("PPPP :%+v\n", string(a))
			//jsonByte, err := ioutil.ReadAll(c.Request.Body)
			//if err != nil {
			//	panic(err.Error())
			//}
			c.Set("payload", payload)
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(ProcessPayload())

	r.GET("/ping", controllers.PingController.Pong)
	r.POST("/divide", controllers.PingController.Divide)
	r.Run()
}
