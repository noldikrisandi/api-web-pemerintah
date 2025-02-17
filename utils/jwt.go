package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var SECRET_KEY = []byte("your-secret-key")

func GenerateJWT(userID string, isAdmin bool) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"is_admin": isAdmin,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SECRET_KEY)
}
