package middlewares

import (
	"konsultanku-app/database/functions"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StudentOnly(c *gin.Context) {

	studentID, err := c.Cookie("UID")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unatuhorized", "data": nil})
		return
	}

	thisStudent := functions.MseOnly(studentID)
	if !thisStudent {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unatuhorized", "data": nil})
		return
	}
	c.Next()
}
