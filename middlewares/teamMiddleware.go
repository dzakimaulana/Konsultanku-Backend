package middlewares

import (
	"konsultanku-app/database/functions"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TeamField(c *gin.Context) {

	teamID, err := c.Cookie("TID")
	if err != nil || teamID == "" {
		uid, _ := c.Cookie("UID")
		teamID, err := functions.InTeam(uid)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "data": nil})
			return
		}
		if teamID != "" {
			c.SetCookie("TID", teamID, 3600, "/", "localhost", false, true)
			c.Next()
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "data": nil})
		return
	}
	c.Next()
}
