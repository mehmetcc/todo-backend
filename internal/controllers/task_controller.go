package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mehmetcc/todo-backend/internal/models"
)

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	db.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var task models.Task
	if err := db.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	db.Create(&task)
	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var task models.Task
	if err := db.First(&task, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	db.Save(&task)
	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var task models.Task
	if err := db.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	db.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
