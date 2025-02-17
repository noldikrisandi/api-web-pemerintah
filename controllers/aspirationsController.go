package controllers

import (
	"api-web-pemerintah/database"
	"api-web-pemerintah/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllAspirations(c *gin.Context) {
	var aspirations []models.Aspirations
	if err := database.DB.Find(&aspirations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data aspirasi", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data aspirasi berhasil diambil", "data": aspirations})
}

func GetAspirationByID(c *gin.Context) {
	id := c.Param("id")

	var aspiration models.Aspirations
	if err := database.DB.Where("id = ?", id).First(&aspiration).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aspirasi tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data aspirasi ditemukan", "data": aspiration})
}

func CreateAspiration(c *gin.Context) {
	var aspiration models.Aspirations

	if err := c.ShouldBindJSON(&aspiration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid", "details": err.Error()})
		return
	}

	idPengirim := aspiration.IdPengirim

	if idPengirim == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Pengirim tidak boleh kosong"})
		return
	}

	sevenDaysAgo := time.Now().AddDate(0, 0, -7)

	var count int64
	if err := database.DB.Model(&models.Aspirations{}).
		Where("id_pengirim = ? AND created_at >= ?", idPengirim, sevenDaysAgo).
		Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghitung aspirasi", "details": err.Error()})
		return
	}

	if count >= 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Anda sudah mengirimkan 2 aspirasi dalam 7 hari terakhir"})
		return
	}

	if err := database.DB.Create(&aspiration).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan aspirasi", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Aspirasi berhasil ditambahkan", "data": aspiration, "total_aspirations_last_week": count})
}

func UpdateAspiration(c *gin.Context) {
	id := c.Param("id")

	var aspiration models.Aspirations
	if err := database.DB.Where("id = ?", id).First(&aspiration).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aspirasi tidak ditemukan"})
		return
	}

	var updatedData models.Aspirations
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid", "details": err.Error()})
		return
	}

	if err := database.DB.Model(&aspiration).Updates(updatedData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui aspirasi", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Aspirasi berhasil diperbarui", "data": aspiration})
}

func DeleteAspiration(c *gin.Context) {
	id := c.Param("id")

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

func GetAspirationsByUserID(c *gin.Context) {
	userID := c.Param("user_id")

	var aspirations []models.Aspirations
	if err := database.DB.Where("id_pengirim = ?", userID).Find(&aspirations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data aspirasi", "details": err.Error()})
		return
	}

	aspirationCount := len(aspirations)

	c.JSON(http.StatusOK, gin.H{
		"message":         "Data aspirasi pengguna ditemukan",
		"data":            aspirations,
		"aspirationCount": aspirationCount,
	})
}
