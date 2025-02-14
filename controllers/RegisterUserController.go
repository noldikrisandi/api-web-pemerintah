package controllers

import (
	"backend/database"
	"backend/models"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user models.Users

	// Bind JSON ke struct
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format data tidak valid"})
		return
	}

	// Hash password sebelum disimpan
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengenkripsi password"})
		return
	}
	user.Password = hashedPassword

	// Simpan user ke database
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registrasi berhasil!"})
}
