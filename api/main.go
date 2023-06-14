package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	gp := r.Group("/user")
	gp.GET("/:username", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Message": "Pong",
		})
	})
	gp.GET("/name", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Message": "name",
		})
	})
	r.Run(":8000")
}
