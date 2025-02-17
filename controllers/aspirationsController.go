package controllers

import (
	"api-web-pemerintah/database"
	"api-web-pemerintah/models"
	"net/http"
	"time"

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

// Get aspirations count by user ID for the current week
// GetAspirationsCountByUserIDThisWeek - Mengambil jumlah aspirasi yang dikirim oleh user dalam minggu ini
func GetAspirationsCountByUserIDThisWeek(c *gin.Context) {
	idPengirim := c.Param("id_pengirim") // Ambil id_pengirim dari parameter URL

	// Ambil tanggal saat ini
	currentTime := time.Now()

	// Hitung awal minggu ini (Senin)
	startOfWeek := currentTime.AddDate(0, 0, -int(currentTime.Weekday())+1) // Memastikan Senin sebagai awal minggu
	endOfWeek := startOfWeek.Add(7 * 24 * time.Hour)                        // Akhir minggu (Minggu depan)

	// Hitung jumlah aspirasi yang dikirim oleh user dalam minggu ini
	var count int64
	if err := database.DB.Model(&models.Aspirations{}).Where("id_pengirim = ? AND created_at BETWEEN ? AND ?", idPengirim, startOfWeek, endOfWeek).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghitung jumlah aspirasi", "details": err.Error()})
		return
	}

	// Mengembalikan response dengan jumlah aspirasi
	c.JSON(http.StatusOK, gin.H{
		"message": "Jumlah aspirasi dalam minggu ini berhasil diambil",
		"count":   count,
	})
}

func CreateAspiration(c *gin.Context) {
	var aspiration models.Aspirations

	// Validasi input JSON
	if err := c.ShouldBindJSON(&aspiration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid", "details": err.Error()})
		return
	}

	// Ambil ID user (id_pengirim) dari input JSON
	idPengirim := aspiration.IdPengirim

	// Pastikan id_pengirim ada
	if idPengirim == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Pengirim tidak boleh kosong"})
		return
	}

	// Simpan aspirasi baru ke database
	if err := database.DB.Create(&aspiration).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan aspirasi", "details": err.Error()})
		return
	}

	// Mengembalikan response dengan status Created dan data aspirasi yang baru
	c.JSON(http.StatusCreated, gin.H{"message": "Aspirasi berhasil ditambahkan", "data": aspiration})
}

// CreateAspiration2 untuk menambahkan aspirasi baru dengan pembatasan 2 aspirasi per minggu
func CreateAspiration2(c *gin.Context) {
	var aspiration models.Aspirations

	// Validasi input JSON
	if err := c.ShouldBindJSON(&aspiration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid", "details": err.Error()})
		return
	}

	// Ambil ID user (id_pengirim) dari input JSON
	idPengirim := aspiration.IdPengirim

	// Pastikan id_pengirim ada
	if idPengirim == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Pengirim tidak boleh kosong"})
		return
	}

	// Tentukan rentang waktu 7 hari terakhir
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)

	// Hitung jumlah aspirasi yang sudah dikirim oleh user dalam 7 hari terakhir
	var count int64
	if err := database.DB.Model(&models.Aspirations{}).
		Where("id_pengirim = ? AND created_at >= ?", idPengirim, sevenDaysAgo).
		Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghitung aspirasi", "details": err.Error()})
		return
	}

	// Cek apakah jumlah aspirasi lebih dari 2
	if count >= 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Anda sudah mengirimkan 2 aspirasi dalam 7 hari terakhir"})
		return
	}

	// Simpan aspirasi baru ke database
	if err := database.DB.Create(&aspiration).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan aspirasi", "details": err.Error()})
		return
	}

	// Mengembalikan response dengan status Created dan data aspirasi yang baru
	c.JSON(http.StatusCreated, gin.H{"message": "Aspirasi berhasil ditambahkan", "data": aspiration, "total_aspirations_last_week": count})
}

// Update aspiration by ID
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

// Get aspirations by user ID
func GetAspirationsByUserID(c *gin.Context) {
	userID := c.Param("user_id") // Ambil user_id dari parameter URL

	var aspirations []models.Aspirations
	if err := database.DB.Where("id_pengirim = ?", userID).Find(&aspirations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data aspirasi", "details": err.Error()})
		return
	}

	// Hitung jumlah aspirasi yang dikirim oleh user
	aspirationCount := len(aspirations)

	c.JSON(http.StatusOK, gin.H{
		"message":         "Data aspirasi pengguna ditemukan",
		"data":            aspirations,
		"aspirationCount": aspirationCount, // Menampilkan jumlah aspirasi
	})
}
