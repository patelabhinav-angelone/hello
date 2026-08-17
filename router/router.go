package router

import (
	"hello/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server Health is good",
		})
	})

	auth := r.Group("/api/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
	}

	p := r.Group("/api")
	{
		p.GET("/users", handler.GetAllUsers)
		p.GET("/users/:id", handler.GetUserByID)
		p.POST("/users", handler.CreateUser)
	}
	return r
}
