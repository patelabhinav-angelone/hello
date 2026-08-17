package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"hello/config"
	"hello/models"
)

func CreateTweet(c *gin.Context) {
	var input models.CreateTweet

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input",
			"details": err.Error(),
		})
		return
	}

	if strings.TrimSpace(input.Content) == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input",
			"details": "content cannot be empty or contain only whitespace",
		})
		return
	}

	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unable to determine user"})
		return
	}

	tweet := models.Tweet{
		Content: input.Content,
		UserID:  userID.(uint),
	}
	res := config.DB.Create(&tweet)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create tweet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "tweet created successfully",
		"data":    tweet,
	})
}

func GetAllTweets(c *gin.Context) {
	var tweets []models.Tweet
	res := config.DB.Order("created_at desc").Find(&tweets)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch tweets"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": tweets})
}

func GetTweetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tweet id"})
		return
	}
	var tweet models.Tweet
	res := config.DB.First(&tweet, id)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tweet not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": tweet})
}
