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
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
        
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

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	tweet := models.Tweet{
		Content: input.Content,
		UserID:  uint(userID),
	}
	res := config.DB.Create(&tweet)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create tweet"})
		return
	}
	tweet.Username = user.Name

	c.JSON(http.StatusOK, gin.H{
		"message": "tweet created successfully",
		"data":    tweet,
	})
}

func GetAllTweets(c *gin.Context) {
	var tweets []models.Tweet
	res := config.DB.Preload("User").Order("created_at desc").Find(&tweets)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch tweets"})
		return
	}
	for i := range tweets {
		tweets[i].Username = tweets[i].User.Name
	}
	c.JSON(http.StatusOK, gin.H{
		"data": tweets})
}

func GetTweetsByUserID(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var tweets []models.Tweet
	res := config.DB.Preload("User").Where("user_id = ?", userID).Order("created_at desc").Find(&tweets)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch tweets"})
		return
	}
	for i := range tweets {
		tweets[i].Username = tweets[i].User.Name
	}

	c.JSON(http.StatusOK, gin.H{
		"data": tweets,
	})
}

func UpdateTweet(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	tweetID, err := strconv.ParseUint(c.Param("tweetId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tweet id"})
		return
	}

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

	var tweet models.Tweet
	res := config.DB.Preload("User").First(&tweet, tweetID)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tweet not found"})
		return
	}

	if tweet.UserID != uint(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not authorized to update this tweet"})
		return
	}

	tweet.Content = input.Content
	res = config.DB.Save(&tweet)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to update tweet"})
		return
	}
	tweet.Username = tweet.User.Name

	c.JSON(http.StatusOK, gin.H{
		"message": "tweet updated successfully",
		"data":    tweet,
	})
}

func DeleteTweet(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	tweetID, err := strconv.ParseUint(c.Param("tweetId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tweet id"})
		return
	}

	var tweet models.Tweet
	res := config.DB.First(&tweet, tweetID)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tweet not found"})
		return
	}

	if tweet.UserID != uint(userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not authorized to delete this tweet"})
		return
	}

	res = config.DB.Delete(&tweet)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to delete tweet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "tweet deleted successfully",
	})
}
