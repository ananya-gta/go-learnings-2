package main

import (
	"net/http"

	"example.com/building-a-rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/events", getEvents)
	router.POST("/create-event", createEvent)

	router.Run(":8080")
}

func getEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event

	err := c.ShouldBindJSON(&event) // maps the JSON response data directly to the struct

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()

	c.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}
