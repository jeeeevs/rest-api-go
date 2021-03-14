package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jeeeevs/timeutil"
)

var upTimeNow = timeutil.UpTime()

func healthcheck(c *gin.Context) {
	upTime := upTimeNow()
	c.JSON(200, gin.H{
		"status": "available",
		"upTime": upTime,
	})
}
func main() {
	// fancy.Hello()
	router := gin.Default()
	router.GET("/healthcheck", healthcheck)
	router.Run(":3000")
}
