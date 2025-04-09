package routes

import (
	"github.com/gin-gonic/gin"
	"schoolwebsite/controllers"
)

func StudentRoutes(r *gin.Engine) {
	r.GET("/students", controllers.GetStudents)
	r.POST("/students", controllers.CreateStudent)
	r.PUT("/students/:id", controllers.UpdateStudent)
	r.DELETE("/students/:student_id/teachers/:teacher_id", controllers.DeleteTeacherFromStudent)

}
