package controllers

import (
	"backend/database"
	"backend/models"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login Admin
func LoginAdminController(c *gin.Context) {
	var loginReq struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Ambil data dari request JSON
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format request salah"})
		return
	}

	var admin models.Admins
	// Cari admin berdasarkan email
	if err := database.DB.Where("email = ?", loginReq.Email).First(&admin).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email tidak ditemukan"})
		return
	}

	// Cek apakah password cocok
	if !utils.CheckPasswordHash(loginReq.Password, admin.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	// Jika password benar, buat token JWT
	token, err := utils.GenerateToken(admin.ID, admin.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat token"})
		return
	}

	// Beri response dengan token
	c.JSON(http.StatusOK, gin.H{"token": token})
}
