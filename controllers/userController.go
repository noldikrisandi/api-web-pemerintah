package controllers

import (
	"api-web-pemerintah/database"
	"api-web-pemerintah/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ambil semua data Users
func GetAllUsers(c *gin.Context) {
	var users []models.Users
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data users", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Ambil satu user berdasarkan ID
func GetUsersByID(c *gin.Context) {
	id := c.Param("id")
	var users models.Users
	if err := database.DB.Where("id = ?", id).First(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Buat user baru
func CreateUsers(c *gin.Context) {
	var users models.Users
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid", "details": err.Error()})
		return
	}

	// Validasi data
	if users.Email == "" || users.Nama == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email dan Nama tidak boleh kosong"})
		return
	}

	if err := database.DB.Create(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data", "details": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, users)
}

// Update user berdasarkan ID
func UpdateUsers(c *gin.Context) {
	id := c.Param("id")
	var users models.Users
	if err := database.DB.Where("id = ?", id).First(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan", "details": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid", "details": err.Error()})
		return
	}

	if err := database.DB.Save(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui data", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Hapus user berdasarkan ID
func DeleteUsers(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Where("id = ?", id).Delete(&models.Users{}).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User berhasil dihapus"})
}
