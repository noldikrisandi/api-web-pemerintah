package middleware

import (
	"api-web-pemerintah/utils" // Mengimpor package utils, bukan middleware
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware adalah middleware untuk memeriksa token JWT di header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Mengambil token dari header Authorization (Bearer token)
		token := authHeader[len("Bearer "):]
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}

		// Verifikasi token menggunakan fungsi dari utils
		claims, err := utils.VerifyJWTToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Jika token valid, simpan userID di konteks untuk digunakan di route selanjutnya
		c.Set("userID", claims.ID)
		c.Set("email", claims.Email)

		c.Next()
	}
}
