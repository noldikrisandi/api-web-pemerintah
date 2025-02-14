package controllers

import (
	"backend/database"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ambil semua data Admins
func GetAllAdmins(c *gin.Context) {
	var admins []models.Admins
	if err := database.DB.Find(&admins).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, admins)
}

// Ambil satu berita berdasarkan ID
func GetAdminsByID(c *gin.Context) {
	id := c.Param("id")
	var admins models.Admins
	if err := database.DB.Where("id = ?", id).First(&admins).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admins not found"})
		return
	}
	c.JSON(http.StatusOK, admins)
}

// Buat berita baru
func CreateAdmins(c *gin.Context) {
	var admins models.Admins
	if err := c.ShouldBindJSON(&admins); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&admins).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, admins)
}

// Update berita berdasarkan ID
func UpdateAdmins(c *gin.Context) {
	id := c.Param("id")
	var admins models.Admins
	if err := database.DB.Where("id = ?", id).First(&admins).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admins not found"})
		return
	}

	if err := c.ShouldBindJSON(&admins); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&admins).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, admins)
}

// Hapus berita berdasarkan ID
func DeleteAdmins(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Where("id = ?", id).Delete(&models.Admins{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admins not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Admins deleted"})
}
