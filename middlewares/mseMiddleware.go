package middlewares

import (
	"konsultanku-app/database/functions"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MseField(c *gin.Context) {

	mseID := NeedUserID(c)
	if mseID == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unatuhorized", "data": nil})
		return
	}
	c.SetCookie("UID", mseID, 3600, "/", "localhost", false, true)
	thisMse := functions.MseRole(mseID)
	if !thisMse {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unatuhorized", "data": nil})
		return
	}
	c.Next()
}
