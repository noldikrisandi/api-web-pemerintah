package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claim struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func VerifyJWTToken(tokenString string) (*Claim, error) {
	Claim := &Claim{}

	token, err := jwt.ParseWithClaims(tokenString, Claim, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("metode signing tidak sesuai: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("token tidak valid: %v", err)
	}

	if Claim.ExpiresAt.Time.Before(time.Now()) {
		return nil, fmt.Errorf("token sudah expired")
	}

	return Claim, nil
}
