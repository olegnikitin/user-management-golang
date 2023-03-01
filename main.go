package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "user was fetched",
		})
	})
	r.POST("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "user was created",
		})
	})
	r.PUT("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "user was created",
		})
	})
	r.DELETE("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "user was deleted",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
