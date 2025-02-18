package controllers

import (
	"api-web-pemerintah/database"
	"api-web-pemerintah/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAdminControllers(c *gin.Context) {
	var adminControllers []models.Admincontroller

	if err := database.DB.Find(&adminControllers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, adminControllers)
}

func CreateAdminController(c *gin.Context) {
	var adminController models.Admincontroller
	if err := c.ShouldBindJSON(&adminController); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var admin models.Admins
	if err := database.DB.Where("id = ?", adminController.IDAdmin).First(&admin).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}

	var aspirasi models.Aspirations
	if err := database.DB.Where("id = ?", adminController.IDAspirasi).First(&aspirasi).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aspirasi not found"})
		return
	}

	if err := database.DB.Create(&adminController).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Admin assigned to aspiration successfully", "data": adminController})
}

func DeleteAdminController(c *gin.Context) {
	idAdmin := c.Param("idadmin")
	idAspirasi := c.Param("idaspirasi")

	var adminController models.Admincontroller
	if err := database.DB.Where("idadmin = ? AND idaspirasi = ?", idAdmin, idAspirasi).First(&adminController).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	if err := database.DB.Delete(&adminController).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}

func DistributeAspirasi1(c *gin.Context) {
	var admins []models.Admins
	if err := database.DB.Find(&admins).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data admin", "details": err.Error()})
		return
	}

	var aspirasi []models.Aspirations
	if err := database.DB.Where("status = ?", "belum ditangani").Find(&aspirasi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data aspirasi", "details": err.Error()})
		return
	}

	var adminWorkloads = make(map[string]int)
	for _, admin := range admins {
		var count int64
		database.DB.Model(&models.Admincontroller{}).Where("idadmin = ?", admin.ID).Count(&count)
		adminWorkloads[admin.ID] = int(count)
	}

	for _, asp := range aspirasi {
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
		adminController.IDAspirasi = asp.ID

		if err := database.DB.Create(&adminController).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendistribusikan aspirasi", "details": err.Error()})
			return
		}

		adminWorkloads[leastLoadedAdminID]++
	}

	c.JSON(http.StatusOK, gin.H{"message": "Aspirasi berhasil didistribusikan ke admin yang sesuai"})
}

func GetAspirationsByAdminID(c *gin.Context) {
	idAdmin := c.Param("id")
	fmt.Println("ID Admin:", idAdmin)

	var adminControllers []models.Admincontroller
	if err := database.DB.Where("idadmin = ?", idAdmin).Find(&adminControllers).Error; err != nil {
		fmt.Println("Error fetching admincontroller:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data admincontroller", "details": err.Error()})
		return
	}

	if len(adminControllers) == 0 {
		fmt.Println("No admincontrollers found")
		c.JSON(http.StatusNotFound, gin.H{"message": "Tidak ada aspirasi untuk admin ini"})
		return
	}

	var aspirationIDs []string
	for _, adminController := range adminControllers {
		aspirationIDs = append(aspirationIDs, adminController.IDAspirasi)
	}

	fmt.Println("Aspiration IDs:", aspirationIDs)

	if len(aspirationIDs) > 0 {
		var aspirations []models.Aspirations
		query := "SELECT * FROM aspirations WHERE id = ANY($1)"
		if err := database.DB.Raw(query, aspirationIDs).Scan(&aspirations).Error; err != nil {
			fmt.Println("Error fetching aspirations:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data aspirasi", "details": err.Error()})
			return
		}

		if len(aspirations) == 0 {
			fmt.Println("No aspirations found")
		}

		c.JSON(http.StatusOK, gin.H{"data": aspirations})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Tidak ada aspirasi untuk admin ini"})
	}
}
