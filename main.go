package main

import (
    "net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	name  string
}
  
func main() {
    dbURL := "postgres://postgres:123456@localhost:5432/gorm"
    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
  
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/users", func(c *gin.Context) {
		var user User

		c.IndentedJSON(http.StatusOK, db.Find(user))
	})
	r.POST("/users", func(c *gin.Context) {
		var user User

		if err := c.BindJSON(&user); err != nil {
			return
		}

		db.Model(user).Create(user)
		c.IndentedJSON(http.StatusCreated, user)
	})
	r.PUT("/users", func(c *gin.Context) {
		var user User

		if err := c.BindJSON(&user); err != nil {
			return
		}

		db.Model(user).Updates(user)
		c.IndentedJSON(http.StatusAccepted, user)
	})
	r.DELETE("/users", func(c *gin.Context) {
		var user User
		id := c.Param("id")

		db.Delete(user, id)
		c.JSON(200, gin.H{
			"message": "user was deleted",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
