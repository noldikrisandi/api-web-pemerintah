package controllers

import (
	"api-web-pemerintah/database"
	"api-web-pemerintah/models"
	"api-web-pemerintah/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginAdminController(c *gin.Context) {
	var loginReq struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format request salah"})
		return
	}

	var admin models.Admins
	if err := database.DB.Where("email = ?", loginReq.Email).First(&admin).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email tidak ditemukan"})
		return
	}

	if !utils.CheckPasswordHash(loginReq.Password, admin.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	token, err := utils.GenerateToken(admin.ID, admin.Email, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func RegisterAdminController(c *gin.Context) {
	var adminReq struct {
		ID       string `json:"id" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		Answer   string `json:"answer" binding:"required"`
	}

	if err := c.ShouldBindJSON(&adminReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format request salah"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminReq.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengenkripsi password"})
		return
	}

	newAdmin := models.Admins{
		ID:       adminReq.ID,
		Email:    adminReq.Email,
		Password: string(hashedPassword),
		Answer:   adminReq.Answer,
	}

	if err := database.DB.Create(&newAdmin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan admin"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Admin berhasil didaftarkan", "id": newAdmin.ID})
}
