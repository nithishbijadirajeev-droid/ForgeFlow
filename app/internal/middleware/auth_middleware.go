package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	jwtutil "github.com/nithishbijadirajeev-droid/forgeflow/app/pkg/jwt"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing authorization header",
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwtutil.ValidateToken(tokenString)

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			return
		}

		c.Next()
	}
}