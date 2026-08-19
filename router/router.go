package router

import (
	"hello/handler"
	"hello/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

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
		p.GET("/users/search", handler.SearchUsers)
		p.GET("/users/:id", handler.GetUserByID)
		p.GET("/users/me", middleware.AuthMiddleware(), handler.GetCurrentUser)
		p.PUT("/users/:id/tweets/:tweetId", middleware.AuthMiddleware(), handler.UpdateTweet)
		p.DELETE("/users/:id/tweets/:tweetId", middleware.AuthMiddleware(), handler.DeleteTweet)
	}

	tweets := r.Group("/api/tweets")
	{
		tweets.GET("", handler.GetAllTweets)
		tweets.GET("/user/:id", handler.GetTweetsByUserID)
		tweets.POST("/user/:id", middleware.AuthMiddleware(), handler.CreateTweet)
	}
	return r
}
