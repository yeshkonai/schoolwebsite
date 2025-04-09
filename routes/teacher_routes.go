package routes

import (
	"github.com/gin-gonic/gin"
	"schoolwebsite/controllers"
)

func RegisterTeacherRoutes(r *gin.Engine) {
	r.GET("/teachers", controllers.GetTeachers)
	r.POST("/teachers", controllers.CreateTeacher)
	r.PUT("/teachers/:id", controllers.UpdateTeacher)
	r.DELETE("/teachers/:id", controllers.DeleteTeacher)
}
