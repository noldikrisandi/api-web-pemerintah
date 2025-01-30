package controllers

import (
	"backend/database"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all subsidies
func GetAllSubsidies(c *gin.Context) {
	var subsidies []models.Subsidies
	if err := database.DB.Find(&subsidies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subsidies)
}

// Get subsidy by ID
func GetSubsidyByID(c *gin.Context) {
	id := c.Param("id")
	var subsidy models.Subsidies
	if err := database.DB.Where("id = ?", id).First(&subsidy).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subsidy not found"})
		return
	}
	c.JSON(http.StatusOK, subsidy)
}

// Create new subsidy
func CreateSubsidy(c *gin.Context) {
	var subsidy models.Subsidies
	if err := c.ShouldBindJSON(&subsidy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&subsidy).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, subsidy)
}

// Update subsidy by ID
func UpdateSubsidy(c *gin.Context) {
	id := c.Param("id")
	var subsidy models.Subsidies
	if err := database.DB.Where("id = ?", id).First(&subsidy).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subsidy not found"})
		return
	}

	if err := c.ShouldBindJSON(&subsidy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&subsidy).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subsidy)
}

// Delete subsidy by ID
func DeleteSubsidy(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Where("id = ?", id).Delete(&models.Subsidies{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subsidy not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subsidy deleted"})
}
