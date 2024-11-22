package routes

import (
	"net/http"
	"strconv"

	"example.com/building-a-rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events", "error": err.Error()})
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

func createEvents(c *gin.Context) {
	// Define a slice to hold multiple events
	var events []models.Event

	// Bind the incoming JSON to the events slice
	err := c.ShouldBindJSON(&events)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	// Iterate over the events and save them one by one
	for _, event := range events {
		// Set default values for ID and UserID if needed
		event.UserID = 1 // Assuming a default user ID for now
		err = event.Save()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
			return
		}
	}

	// Return a response indicating that all events have been created
	c.JSON(http.StatusCreated, gin.H{"message": "events created", "events": events})
}


func getEventByID(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id", "error": err.Error()})
		return
	}

	event, err := models.GetEventByID(eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not fetch event id", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event successfully fetched: ", "event": event})
}

func updateEventByID(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id", "error": err.Error()})
		return
	}
	_, err = models.GetEventByID(eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not fetch event id", "error": err.Error()})
		return
	}

	var updatedEvent models.Event

	err = c.ShouldBindJSON(&updatedEvent) // maps the JSON response data directly to the struct

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	updatedEvent.ID = eventID
	err = updatedEvent.UpdateEventByID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event successfully updated: ", "event": updatedEvent})
}

func deleteEvent(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id", "error": err.Error()})
		return
	}
	
	event, err := models.GetEventByID(eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not fetch event id", "error": err.Error()})
		return
	}

	err = event.DeleteEvent()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event successfully deleted!"})
}