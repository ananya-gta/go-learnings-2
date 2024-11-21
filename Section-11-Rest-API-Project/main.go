package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.GET("/events", )

	server.Run(":8080")
}

func getEvents(c *gin.Context) {
	
}
