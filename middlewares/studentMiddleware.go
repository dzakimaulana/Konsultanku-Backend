package middlewares

import (
	"konsultanku-app/database/functions"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StudentField(c *gin.Context) {

	studentID := NeedUserID(c)
	if studentID == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unatuhorized", "data": nil})
		return
	}
	c.SetCookie("UID", studentID, 3600, "/", "localhost", false, true)
	thisStudent := functions.StudentRole(studentID)
	if !thisStudent {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unatuhorized", "data": nil})
		return
	}
	c.Next()
}
