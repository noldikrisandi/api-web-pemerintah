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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data aspirasi", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data aspirasi berhasil diambil", "data": aspirations})
}

// Get aspiration by ID
func GetAspirationByID(c *gin.Context) {
	id := c.Param("id") // Ambil ID sebagai string

	var aspiration models.Aspirations
	if err := database.DB.Where("id = ?", id).First(&aspiration).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aspirasi tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data aspirasi ditemukan", "data": aspiration})
}

// Create new aspiration
func CreateAspiration(c *gin.Context) {
	var aspiration models.Aspirations

	// Validasi input JSON
	if err := c.ShouldBindJSON(&aspiration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid", "details": err.Error()})
		return
	}

	// Simpan ke database
	if err := database.DB.Create(&aspiration).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan aspirasi", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Aspirasi berhasil ditambahkan", "data": aspiration})
}

// Update aspiration by ID
// Update all fields of the aspiration by ID
func UpdateAspiration(c *gin.Context) {
	id := c.Param("id") // Ambil ID sebagai string

	var aspiration models.Aspirations
	// Ambil data aspirasi berdasarkan ID
	if err := database.DB.Where("id = ?", id).First(&aspiration).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aspirasi tidak ditemukan"})
		return
	}

	// Menangani input dari request body
	var updatedData models.Aspirations
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid", "details": err.Error()})
		return
	}

	// Memperbarui seluruh data aspirasi
	if err := database.DB.Model(&aspiration).Updates(updatedData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui aspirasi", "details": err.Error()})
		return
	}

	// Mengembalikan response sukses
	c.JSON(http.StatusOK, gin.H{"message": "Aspirasi berhasil diperbarui", "data": aspiration})
}

// Delete aspiration by ID
func DeleteAspiration(c *gin.Context) {
	id := c.Param("id") // Ambil ID sebagai string

	var aspiration models.Aspirations
	if err := database.DB.Where("id = ?", id).First(&aspiration).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aspirasi tidak ditemukan"})
		return
	}

	if err := database.DB.Delete(&aspiration).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus aspirasi", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Aspirasi berhasil dihapus"})
}
