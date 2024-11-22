package main

import (
	"example.com/building-a-rest-api/db"
	"github.com/gin-gonic/gin"
	"example.com/building-a-rest-api/routes"
)

func main() {

	db.InitDB()
	router := gin.Default()

	routes.RegisterRoutes(router)
	

	router.Run(":8080")
}
