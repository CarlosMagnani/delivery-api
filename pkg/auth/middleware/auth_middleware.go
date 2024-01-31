package middleware

import (
	"delivery_api/pkg/auth/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService *service.AuthJWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is must to use this route"})
			return
		}

		isValid, err := authService.VerifyToken(token)
		if err != nil || !isValid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
			return
		}

		c.Next()
	}
}