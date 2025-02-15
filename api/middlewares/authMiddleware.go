package middlewares

import (
	"cvwo-backend/api/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is used to protect routes with JWT authentication
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		tokenString, err := c.Cookie("jwt")
		if(err != nil){
			log.Print(err)
		}
		
		
		// Verify the token
		claims, err := utils.VerifyJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}
		// Attach the claims to the context for further use
		c.Set("userId", claims["userId"])
		c.Next()
	}
}

func GetUserIdMiddleware() gin.HandlerFunc {
	return func(c*gin.Context) {
		tokenString, err := c.Cookie("jwt")
		if(err != nil){
			log.Print(err)
		}
		// Verify the token
		claims, err := utils.VerifyJWT(tokenString)
		if err != nil {
			c.Next()
		}
		c.Set("userId", claims["userId"])
		c.Next()
	}
}

