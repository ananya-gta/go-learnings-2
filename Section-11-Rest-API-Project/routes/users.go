package routes

import (
	"net/http"
	// "strconv"

	"example.com/building-a-rest-api/models"
	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data", "error": err.Error()})
		return
	}

	err = user.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not signup user. Try again later."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created", "user": user})
}