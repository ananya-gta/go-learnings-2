package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/events", getEvents)

	router.Run(":8080")
}

func getEvents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message":"Hello"})
}