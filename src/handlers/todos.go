package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gorm.io/gorm"
	"strconv"
)

type Todo struct {
	ID  int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func GetTodos(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todos []Todo
		if err := db.Find(&todos).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, todos)
	}
}

func CreateTodo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newTodo Todo
		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&newTodo).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, newTodo)
	}
}

func UpdateTodo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		var updated Todo
		if err := c.ShouldBindJSON(&updated); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Model(&Todo{}).Where("id = ?", id).Updates(updated).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		updated.ID = id // Ensure the ID is set in the response
		c.JSON(http.StatusOK, updated)
	}
}

func DeleteTodo(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		if err := db.Delete(&Todo{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"deleted": id})
	}
}