package routes

import (
	"apiv2/pkg/models"
	"apiv2/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// signup users functionality/endpoint
func CreateUsers(c *gin.Context) {
	var user models.User

	// Parse the incoming JSON body into the users struct
	if err := c.ShouldBindJSON(&user); err != nil {
		// Return a bad request error with specific message
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = 1

	// // Validate the users data
	// if err := users.Checkusers(); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// Save the users to the database
	if err := user.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message and the created users details
	c.JSON(http.StatusCreated, gin.H{
		"message": "user '" + user.Name + "' created successfully",
		"users":   user.Email,
	})
}

// login users
func Login(c *gin.Context) {
	var user models.User

	// Parse the incoming JSON body into the users struct
	if err := c.ShouldBindJSON(&user); err != nil {
		// Return a bad request error with specific message
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := user.ValidateCreds()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully", "token": token})
}
