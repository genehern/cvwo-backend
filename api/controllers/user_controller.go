package controllers

import (
	"cvwo-backend/api/models"
	"cvwo-backend/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Insert user into the database
	if err := models.CreateUser(models.DB, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}


func Login(c *gin.Context){
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	
	if err := models.ValidateUser(models.DB, &user); err!=nil { // Placeholder validation
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
		token, err := utils.GenerateJWT(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}
		
		c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully"})
}