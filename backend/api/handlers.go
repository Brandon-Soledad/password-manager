package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// /////////////////////////////////////////////////////////
// ///////////          USER ENDPOINTS         /////////////
// /////////////////////////////////////////////////////////
func CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func GetUser(c *gin.Context) {
	// Get ID from request URL
	id := c.Param("id")

	// Convert ID into integer
	var userID int32
	_, err := fmt.Scanln(id, "%d", &userID)
	if err != nil {
		return
	}

	// Retrieve the user from the database
	var user User
	result := db.First(&user, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Return the user as JSON response
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateUserPassword(c *gin.Context) {
	// Get the ID parameter from the request URL
	id := c.Param("id")

	// Convert the ID to an integer
	var userID int32
	_, err := fmt.Sscanf(id, "%d", &userID)
	if err != nil {
		return
	}

	// Retrieve the user from the database
	var user User
	result := db.First(&user, userID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Bind JSON body into a new struct to extract the password
	var newPassword struct {
		Password string `json:"password"`
	}

	if err := c.BindJSON(&newPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Update the password field
	user.Password = newPassword.Password
	db.Save(&user)

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

func DeleteUser(c *gin.Context) {
	// Get the ID parameter from the request URL
	id := c.Param("id")

	// Convert the ID to an integer
	var userID int32
	_, err := fmt.Sscanf(id, "%d", &userID)
	if err != nil {
		return
	}

	// Delete the user from the database
	result := db.Delete(&User{}, userID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
