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
		teamID, err = functions.InTeam(uid)
		if err != nil || teamID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "data": nil})
			return
		}
	}
	c.SetCookie("TID", teamID, 3600, "/", "localhost", false, true)
	c.Next()
}

func IsLeader(c *gin.Context) {

	studentID, _ := c.Cookie("UID")
	isLeader := functions.IsLeader(studentID)
	if !isLeader {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "data": nil})
		return
	}
	c.Next()
}
