package controllers

import (
	"api-web-pemerintah/database"
	"api-web-pemerintah/models"
	"fmt"
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

func GetAspirationsCountByUserID(c *gin.Context) {
	idPengirim := c.Param("id_pengirim")

	var count int64
	err := database.DB.Model(&models.Aspirations{}).
		Where("id_pengirim = ?", idPengirim).
		Count(&count).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghitung aspirasi", "details": err.Error()})
		return
	}

	fmt.Println("Jumlah aspirasi:", count)

	c.JSON(http.StatusOK, gin.H{"total_aspirations": count})
}

func CreateAspiration(c *gin.Context) {
	var aspiration models.Aspirations

	if err := c.ShouldBindJSON(&aspiration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Format data tidak valid",
			"detail": err.Error(),
		})
		return
	}

	if aspiration.IdPengirim == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Pengirim tidak boleh kosong"})
		return
	}

	idPengirim := aspiration.IdPengirim
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)

	var count int64
	if err := database.DB.Model(&models.Aspirations{}).
		Where("id_pengirim = ? AND created_at >= ?", idPengirim, sevenDaysAgo).
		Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Gagal menghitung aspirasi",
			"details": err.Error(),
		})
		return
	}

	if count >= 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Anda sudah mengirimkan 2 aspirasi dalam 7 hari terakhir"})
		return
	}

	if aspiration.CreatedAt.IsZero() {
		aspiration.CreatedAt = time.Now()
	}

	if err := database.DB.Create(&aspiration).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Gagal menyimpan aspirasi",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":                     "Aspirasi berhasil ditambahkan",
		"data":                        aspiration,
		"total_aspirations_last_week": count + 1,
	})

	DistributeAspirasi(c, aspiration)
}

func DistributeAspirasi(c *gin.Context, aspiration models.Aspirations) {

	var admins []models.Admins
	if err := database.DB.Find(&admins).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data admin", "details": err.Error()})
		return
	}

	var adminWorkloads = make(map[string]int)
	for _, admin := range admins {
		var count int64
		database.DB.Model(&models.Admincontroller{}).Where("idadmin = ?", admin.ID).Count(&count)
		adminWorkloads[admin.ID] = int(count)
	}

	var leastLoadedAdminID string
	leastWorkload := -1
	for adminID, workload := range adminWorkloads {
		if leastWorkload == -1 || workload < leastWorkload {
			leastWorkload = workload
			leastLoadedAdminID = adminID
		}
	}

	var adminController models.Admincontroller
	adminController.IDAdmin = leastLoadedAdminID
	adminController.IDAspirasi = aspiration.ID

	if err := database.DB.Create(&adminController).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendistribusikan aspirasi", "details": err.Error()})
		return
	}

	adminWorkloads[leastLoadedAdminID]++

	c.JSON(http.StatusOK, gin.H{"message": "Aspirasi berhasil didistribusikan ke admin yang sesuai"})
}

func CreateAspiration2(c *gin.Context) { // sementara belum digunakan karna ada fitur tambaahan
	var aspiration models.Aspirations

	if err := c.ShouldBindJSON(&aspiration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Format data tidak valid",
			"detail": err.Error(),
		})
		return
	}

	if aspiration.IdPengirim == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID Pengirim tidak boleh kosong"})
		return
	}

	idPengirim := aspiration.IdPengirim
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)

	var count int64
	if err := database.DB.Model(&models.Aspirations{}).
		Where("id_pengirim = ? AND created_at >= ?", idPengirim, sevenDaysAgo).
		Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Gagal menghitung aspirasi",
			"details": err.Error(),
		})
		return
	}

	if count >= 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Anda sudah mengirimkan 2 aspirasi dalam 7 hari terakhir"})
		return
	}

	if aspiration.CreatedAt.IsZero() {
		aspiration.CreatedAt = time.Now()
	}

	if err := database.DB.Create(&aspiration).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Gagal menyimpan aspirasi",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":                     "Aspirasi berhasil ditambahkan",
		"data":                        aspiration,
		"total_aspirations_last_week": count + 1,
	})
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
