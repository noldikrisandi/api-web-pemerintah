package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func VerifyJWTToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// verifikasi token JWT
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("metode signing tidak sesuai: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("token tidak valid: %v", err)
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, fmt.Errorf("token sudah expired")
	}

	return claims, nil
}
