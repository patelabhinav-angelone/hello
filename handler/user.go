package handler

import (
	"hello/config"
	"hello/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

func CreateUser(c *gin.Context) {
	var input models.CreateUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input",
			"details": err.Error(),
		})
		return
	}

	hashpass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		Name:     input.Name,
		Password: string(hashpass),
		EmailID:  input.EmailID,
		Age:      input.Age,
	}
	res := config.DB.Create(&user)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created successfully",
		"data":    user,
	})
}
