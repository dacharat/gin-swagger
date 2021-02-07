package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware verify auth middleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Request.Header.Get("X-ID")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid token"})
			c.Abort()
			return
		}

		c.Set("ID", id)
		c.Next()
	}
}
