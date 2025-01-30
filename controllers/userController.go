package controllers

import (
	"backend/database"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ambil semua data Users
func GetAllUsers(c *gin.Context) {
	var users []models.Users
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Ambil satu berita berdasarkan ID
func GetUsersByID(c *gin.Context) {
	id := c.Param("id")
	var users models.Users
	if err := database.DB.Where("id = ?", id).First(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Users not found"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Buat berita baru
func CreateUsers(c *gin.Context) {
	var users models.Users
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, users)
}

// Update berita berdasarkan ID
func UpdateUsers(c *gin.Context) {
	id := c.Param("id")
	var users models.Users
	if err := database.DB.Where("id = ?", id).First(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Users not found"})
		return
	}

	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Hapus berita berdasarkan ID
func DeleteUsers(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Where("id = ?", id).Delete(&models.Users{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Users not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Users deleted"})
}
