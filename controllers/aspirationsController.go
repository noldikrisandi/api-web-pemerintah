package controllers

import (
	"backend/database"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all aspirations
func GetAllAspirations(c *gin.Context) {
	var aspirations []models.Aspirations
	if err := database.DB.Find(&aspirations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, aspirations)
}

// Get aspiration by ID
func GetAspirationByID(c *gin.Context) {
	id := c.Param("id")
	var aspiration models.Aspirations
	if err := database.DB.Where("id = ?", id).First(&aspiration).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aspiration not found"})
		return
	}
	c.JSON(http.StatusOK, aspiration)
}

// Create new aspiration
func CreateAspiration(c *gin.Context) {
	var aspiration models.Aspirations
	if err := c.ShouldBindJSON(&aspiration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&aspiration).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, aspiration)
}

// Update aspiration by ID
func UpdateAspiration(c *gin.Context) {
	id := c.Param("id")
	var aspiration models.Aspirations
	if err := database.DB.Where("id = ?", id).First(&aspiration).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aspiration not found"})
		return
	}

	if err := c.ShouldBindJSON(&aspiration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&aspiration).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, aspiration)
}

// Delete aspiration by ID
func DeleteAspiration(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Where("id = ?", id).Delete(&models.Aspirations{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aspiration not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Aspiration deleted"})
}
