package jwt

import (
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var secret = []byte(getSecret())

func getSecret() string {

	if os.Getenv("JWT_SECRET") == "" {
		return "forgeflow-super-secret-key"
	}

	return os.Getenv("JWT_SECRET")
}

func GenerateToken(userID string) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
}