package controllers

import (
	"backend/database"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get all packages
func GetAllPackages(c *gin.Context) {
	var packages []models.Packages
	if err := database.DB.Find(&packages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, packages)
}

// Get package by ID
func GetPackageByID(c *gin.Context) {
	id := c.Param("id")
	var packages models.Packages
	if err := database.DB.Where("id = ?", id).First(&packages).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Package not found"})
		return
	}
	c.JSON(http.StatusOK, packages)
}

// Create new package
func CreatePackage(c *gin.Context) {
	var packages models.Packages
	if err := c.ShouldBindJSON(&packages); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&packages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, packages)
}

// Update package by ID
func UpdatePackage(c *gin.Context) {
	id := c.Param("id")
	var packages models.Packages
	if err := database.DB.Where("id = ?", id).First(&packages).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Package not found"})
		return
	}

	if err := c.ShouldBindJSON(&packages); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&packages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, packages)
}

// Delete package by ID
func DeletePackage(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Where("id = ?", id).Delete(&models.Packages{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Package not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Package deleted"})
}
