package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Struct untuk menyimpan claims JWT
type Claims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// Ambil secret key dari environment variable
var secretKey = []byte(os.Getenv("JWT_SECRET"))

// VerifyJWTToken digunakan untuk memverifikasi token JWT
func VerifyJWTToken(tokenString string) (*Claims, error) {
	// Inisialisasi claims untuk menampung data dari token
	claims := &Claims{}

	// Verifikasi token JWT
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Pastikan algoritma yang digunakan adalah HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("metode signing tidak sesuai: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	// Cek apakah token valid
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("token tidak valid: %v", err)
	}

	// Cek apakah token sudah expired
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, fmt.Errorf("token sudah expired")
	}

	// Kembalikan claims jika token valid
	return claims, nil
}
