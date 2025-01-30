package controllers

import (
	"backend/database"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all testimonies
func GetAllTestimonies(c *gin.Context) {
	var testimonies []models.Testimonies
	if err := database.DB.Find(&testimonies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, testimonies)
}

// Get testimony by ID
func GetTestimonyByID(c *gin.Context) {
	id := c.Param("id")
	var testimony models.Testimonies
	if err := database.DB.Where("id = ?", id).First(&testimony).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Testimony not found"})
		return
	}
	c.JSON(http.StatusOK, testimony)
}

// Create new testimony
func CreateTestimony(c *gin.Context) {
	var testimony models.Testimonies
	if err := c.ShouldBindJSON(&testimony); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&testimony).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, testimony)
}

// Update testimony by ID
func UpdateTestimony(c *gin.Context) {
	id := c.Param("id")
	var testimony models.Testimonies
	if err := database.DB.Where("id = ?", id).First(&testimony).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Testimony not found"})
		return
	}

	if err := c.ShouldBindJSON(&testimony); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&testimony).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, testimony)
}

// Delete testimony by ID
func DeleteTestimony(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Where("id = ?", id).Delete(&models.Testimonies{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Testimony not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Testimony deleted"})
}
