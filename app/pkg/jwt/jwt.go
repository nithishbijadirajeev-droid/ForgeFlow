package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/nithishbijadirajeev-droid/forgeflow/app/internal/config"
)

func GenerateToken(userID string) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.Get("JWT_SECRET")))
}

func ValidateToken(tokenString string) (*jwt.Token, error) {

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		return []byte(config.Get("JWT_SECRET")), nil

	})
}