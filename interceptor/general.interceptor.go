package interceptor

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GeneralInterceptor1 - call this method to add interceptor
func GeneralInterceptor1(c *gin.Context) {
	token := c.Query("token")
	if token == "1234" {
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.Abort()
	}

}
