package main

import (
	"net/http"
	"strconv"

	"example.com/building-a-rest-api/db"
	"example.com/building-a-rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	router := gin.Default()

	router.GET("/events", getEvents)
	router.POST("/create-event", createEvent)
	router.GET("/events/:id", getEventByID)

	router.Run(":8080")
}

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events" ,"error": err.Error()})
	}
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

	err = event.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}

func getEventByID(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id" ,"error": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not fetch event id" ,"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event successfully fetched: ", "event": event})
}
