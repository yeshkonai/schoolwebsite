package main

import (
	"schoolwebsite/config"
	"schoolwebsite/models"
	"schoolwebsite/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	db.AutoMigrate(&models.Student{}, &models.Teacher{})

	r := gin.Default()
	routes.StudentRoutes(r)
	routes.RegisterTeacherRoutes(r)
	
	r.Run()
}
