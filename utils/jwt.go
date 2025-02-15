package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// SECRET_KEY untuk JWT (sebaiknya simpan di ENV)
var SECRET_KEY = []byte("your-secret-key")

// GenerateJWT membuat token JWT untuk user
// GenerateJWT membuat token JWT untuk user atau admin
func GenerateJWT(userID string, isAdmin bool) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"is_admin": isAdmin,                               // Menambahkan role untuk admin
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token berlaku 24 jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SECRET_KEY)
}
