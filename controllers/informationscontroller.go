package controllers

import (
	"backend/database"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ambil semua data informasi
func GetAllInformations(c *gin.Context) {
	var informations []models.Informations
	if err := database.DB.Find(&informations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, informations)
}

// Ambil satu informasi berdasarkan ID
func GetInformationByID(c *gin.Context) {
	id := c.Param("id")
	var information models.Informations
	if err := database.DB.Where("id = ?", id).First(&information).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Information not found"})
		return
	}
	c.JSON(http.StatusOK, information)
}

// Buat informasi baru
func CreateInformation(c *gin.Context) {
	var information models.Informations
	if err := c.ShouldBindJSON(&information); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&information).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, information)
}

// Update informasi berdasarkan ID
func UpdateInformation(c *gin.Context) {
	id := c.Param("id")
	var information models.Informations
	if err := database.DB.Where("id = ?", id).First(&information).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Information not found"})
		return
	}

	if err := c.ShouldBindJSON(&information); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&information).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, information)
}

// Hapus informasi berdasarkan ID
func DeleteInformation(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Where("id = ?", id).Delete(&models.Informations{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Information not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Information deleted"})
}
