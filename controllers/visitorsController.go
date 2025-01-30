package controllers

import (
	"backend/database"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all visitors
func GetAllVisitors(c *gin.Context) {
	var visitors []models.Visitors
	if err := database.DB.Find(&visitors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, visitors)
}

// Get visitor by ID
func GetVisitorByID(c *gin.Context) {
	id := c.Param("id")
	var visitor models.Visitors
	if err := database.DB.Where("id = ?", id).First(&visitor).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Visitor not found"})
		return
	}
	c.JSON(http.StatusOK, visitor)
}

// Create new visitor
func CreateVisitor(c *gin.Context) {
	var visitor models.Visitors
	if err := c.ShouldBindJSON(&visitor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&visitor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, visitor)
}

// Update visitor by ID
func UpdateVisitor(c *gin.Context) {
	id := c.Param("id")
	var visitor models.Visitors
	if err := database.DB.Where("id = ?", id).First(&visitor).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Visitor not found"})
		return
	}

	if err := c.ShouldBindJSON(&visitor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&visitor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, visitor)
}

// Delete visitor by ID
func DeleteVisitor(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Where("id = ?", id).Delete(&models.Visitors{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Visitor not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Visitor deleted"})
}
