package controllers

import (
	"backend/database"
	"backend/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Secret key untuk JWT dari environment variable
var secretKey []byte

func init() {
	secretKey = []byte(os.Getenv("JWT_SECRET"))
	if len(secretKey) == 0 {
		log.Fatal("JWT_SECRET environment variable is not set")
	}
}

// Struct untuk request login
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginHandler untuk menangani autentikasi dari database
func LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var admin models.Admins
	result := database.DB.Where("email = ?", req.Email).First(&admin)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error", "details": result.Error.Error()})
		return
	}

	// Validasi password dengan bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	// Generate token JWT dengan user ID dan email
	claims := jwt.MapClaims{
		"id":    admin.ID,
		"email": admin.Email,
		"exp":   time.Now().Add(24 * time.Hour).Unix(), // Token berlaku 24 jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":   tokenString,
		"message": "Login successful",
	})
}
