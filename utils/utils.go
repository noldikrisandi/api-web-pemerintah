package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Secret key untuk JWT
var jwtKey = []byte("secret_key") // Ganti dengan secret key yang lebih aman

// Struktur payload token JWT
type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

// Fungsi untuk membuat token JWT
func GenerateToken(userID, email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token berlaku 24 jam
	claims := &Claims{
		UserID: userID,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Buat token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Fungsi untuk verifikasi token JWT
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("token tidak valid")
	}

	return claims, nil
}

// Hash password menggunakan bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Cek apakah password sesuai dengan hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
