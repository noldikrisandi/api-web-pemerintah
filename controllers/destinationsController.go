package controllers

import (
	"backend/database"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all destinations
func GetAllDestinations(c *gin.Context) {
	var destinations []models.Destinations
	if err := database.DB.Find(&destinations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, destinations)
}

// Get destination by ID
func GetDestinationByID(c *gin.Context) {
	id := c.Param("id")
	var destination models.Destinations
	if err := database.DB.Where("id = ?", id).First(&destination).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Destination not found"})
		return
	}
	c.JSON(http.StatusOK, destination)
}

// Create new destination
func CreateDestination(c *gin.Context) {
	var destination models.Destinations
	if err := c.ShouldBindJSON(&destination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&destination).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, destination)
}

// Update destination by ID
func UpdateDestination(c *gin.Context) {
	id := c.Param("id")
	var destination models.Destinations
	if err := database.DB.Where("id = ?", id).First(&destination).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Destination not found"})
		return
	}

	if err := c.ShouldBindJSON(&destination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&destination).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, destination)
}

// Delete destination by ID
func DeleteDestination(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Where("id = ?", id).Delete(&models.Destinations{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Destination not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Destination deleted"})
}
