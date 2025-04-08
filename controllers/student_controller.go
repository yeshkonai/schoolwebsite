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

	// Получаем параметры limit и grade из query
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
	config.DB.Create(&student)
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
	config.DB.Save(&student)
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Student{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
}
