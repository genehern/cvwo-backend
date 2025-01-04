package middlewares

import (
	"net/http"
	"cvwo-backend/api/utils"
	"strings"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware is used to protect routes with JWT authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Extract token from "Bearer token"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}
		tokenString := tokenParts[1]

		// Verify the token
		claims, err := utils.VerifyJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Attach the claims to the context for further use
		c.Set("user", claims["username"])

		c.Next()
	}
}
