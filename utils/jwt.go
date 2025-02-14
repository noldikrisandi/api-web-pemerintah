package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// SECRET_KEY untuk JWT (sebaiknya simpan di ENV)
var SECRET_KEY = []byte("your-secret-key")

// GenerateJWT membuat token JWT untuk user
func GenerateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token berlaku 24 jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SECRET_KEY)
}
