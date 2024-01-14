package middlewares

import (
	"konsultanku-app/database/functions"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MseField(c *gin.Context) {

	mseID, err := c.Cookie("UID")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unatuhorized", "data": nil})
		return
	}

	thisMse := functions.MseOnly(mseID)
	if !thisMse {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unatuhorized", "data": nil})
		return
	}
	c.Next()
}
