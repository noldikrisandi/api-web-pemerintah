package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Secret key untuk JWT
var jwtKey = []byte("your-secret-key") // Ganti dengan secret key yang lebih aman

// Struktur payload token JWT
type Claims struct {
	UserID  string `json:"user_id"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"` // Menambahkan status admin
	jwt.StandardClaims
}

// Fungsi untuk membuat token JWT
func GenerateToken(userID, email string, isAdmin bool) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID:  userID,
		Email:   email,
		IsAdmin: isAdmin, // Menyertakan status admin
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
