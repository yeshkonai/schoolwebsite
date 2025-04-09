package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"schoolwebsite/config"
	"schoolwebsite/models"
)

func GetTeachers(c *gin.Context) {
	var teachers []models.Teacher
	config.DB.Preload("Students").Find(&teachers)
	c.JSON(http.StatusOK, teachers)
}

func CreateTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Применяем студентов к учителям
	for i := range teacher.Students {
		var student models.Student
		if err := config.DB.First(&student, teacher.Students[i].ID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Student not found"})
			return
		}
		teacher.Students[i] = student
	}

	// Создаем учителя с учениками
	if err := config.DB.Create(&teacher).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, teacher)
}

func UpdateTeacher(c *gin.Context) {
	id := c.Param("id")
	var teacher models.Teacher
	if err := config.DB.First(&teacher, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
		return
	}

	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Применяем студентов к учителю
	for i := range teacher.Students {
		var student models.Student
		if err := config.DB.First(&student, teacher.Students[i].ID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Student not found"})
			return
		}
		teacher.Students[i] = student
	}

	// Обновляем учителя с учениками
	if err := config.DB.Save(&teacher).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, teacher)
}

func DeleteTeacher(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Teacher{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete teacher"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Teacher deleted"})
}
