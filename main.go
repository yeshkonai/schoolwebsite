package main

import (
	"schoolwebsite/config"
	"schoolwebsite/models"
	"schoolwebsite/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	db.AutoMigrate(&models.Student{})

	r := gin.Default()
	routes.StudentRoutes(r)

	r.Run() // default on :8080
}
