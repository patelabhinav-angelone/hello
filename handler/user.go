package handler

import (
	"hello/config"
	"hello/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User
	res := config.DB.Find(&users)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users})
}

func GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	var user models.User
	res := config.DB.First(&user, id)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user})
}

func SearchUsers(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query parameter 'q' is required"})
		return
	}

	var users []models.User
	pattern := "%" + query + "%"
	res := config.DB.Where("name ILIKE ? OR email_id ILIKE ?", pattern, pattern).Find(&users)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to search users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func GetCurrentUser(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unable to determine user"})
		return
	}

	var user models.User
	res := config.DB.First(&user, userID)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
