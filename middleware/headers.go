package middleware

import (
	"auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HeaderAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("x-API-KEY")
		if apiKey == "" {
			utils.Logger.Println("x-API-KEY header is required")
			c.JSON(http.StatusBadRequest, gin.H{"error": "x-API-KEY header is required"})
			c.Abort()
			return
		}

		c.Next()
	}
}
