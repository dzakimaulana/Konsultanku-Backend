package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {

	if _, err := c.Cookie("token"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "data": nil})
		return
	}

	if _, err := c.Cookie("UID"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "data": nil})
		return
	}

	c.Next()
}
