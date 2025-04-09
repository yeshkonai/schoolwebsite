package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"schoolwebsite/config"
	"schoolwebsite/models"
	"strconv"
)

func GetStudents(c *gin.Context) {
	var students []models.Student

	limitParam := c.Query("limit")
	gradeParam := c.Query("grade")

	// Начинаем с базы запроса
	query := config.DB

	// Если указан grade, фильтруем по нему
	if gradeParam != "" {
		query = query.Where("grade = ?", gradeParam)
	}

	// Если указан limit, применяем его
	if limitParam != "" {
		limit, err := strconv.Atoi(limitParam)
		if err == nil {
			query = query.Limit(limit)
		}
	}

	// Выполняем запрос
	query.Find(&students)

	c.JSON(http.StatusOK, students)
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Применяем студентов к учителям
	for i := range student.Teachers {
		// Это создаст/обновит учителей и присоединит их к студенту
		var teacher models.Teacher
		if err := config.DB.First(&teacher, student.Teachers[i].ID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Teacher not found"})
			return
		}
		student.Teachers[i] = teacher
	}

	// Создаем студента с учителями
	if err := config.DB.Create(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, student)
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	if err := config.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Применяем студентов к учителям
	for i := range student.Teachers {
		var teacher models.Teacher
		if err := config.DB.First(&teacher, student.Teachers[i].ID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Teacher not found"})
			return
		}
		student.Teachers[i] = teacher
	}

	// Обновляем студента с учителями
	if err := config.DB.Save(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, student)
}

func DeleteTeacherFromStudent(c *gin.Context) {
	studentID := c.Param("student_id")
	teacherID := c.Param("teacher_id")

	var student models.Student
	var teacher models.Teacher

	if err := config.DB.First(&student, studentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	if err := config.DB.First(&teacher, teacherID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
		return
	}

	// Удаляем связь
	if err := config.DB.Model(&student).Association("Teachers").Delete(&teacher); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove teacher"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Teacher removed from student"})
}
