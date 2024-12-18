package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/events", getEvents)
	router.POST("/createEvent", createEvent)
	router.POST("/createEvents", createEvents)
	router.GET("/events/:id", getEventByID)
	router.PUT("/updateEvent/:id", updateEventByID)
	router.DELETE("deleteEvent/:id", deleteEvent)
	router.POST("/signup", signup)
	router.POST("/login", login)
}