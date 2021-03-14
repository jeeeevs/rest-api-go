package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jeeeevs/rest-api-go/controller"
)

func main() {
	router := gin.Default()
	router.GET("/healthcheck", func(c *gin.Context) {
		msg := controller.Healthcheck()
		c.JSON(200, msg)
	})
	router.Run(":3000")
}
