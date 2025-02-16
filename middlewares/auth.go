package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

// Middleware untuk autentikasi JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Kamu tidak bisa mengakses API ini tanpa token"})
			c.Abort()
			return
		}

		// Harus dalam format "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Format token tidak valid"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Verifikasi token JWT
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			// Pastikan algoritma yang digunakan HS256
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("metode signing tidak sesuai: %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		// Cek apakah token valid
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token tidak valid"})
			c.Abort()
			return
		}

		// Cek apakah token sudah expired
		if claims.ExpiresAt.Time.Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token sudah expired"})
			c.Abort()
			return
		}

		// Jika token valid, simpan data user di context
		c.Set("userID", claims.ID)
		c.Set("email", claims.Email)

		c.Next()
	}
}
