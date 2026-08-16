package handler

import (
	"hello/config"
	"hello/models"
	"net/http"

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
	id := c.Param("id")
	var user models.User
	res := config.DB.First(&user, id)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user})
}
