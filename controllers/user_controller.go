package controllers

import (
	"go-gin-api/db"
	"go-gin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	// Bind incoming JSON to the User struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save to DB
	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created user
	c.JSON(http.StatusCreated, user)
}

func GetUsers(c *gin.Context) {
	var users []models.User
	db.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}