package middlewares

import (
	"konsultanku-app/controllers"
	"konsultanku-app/database"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {

	token, err := c.Cookie("token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "data": nil})
		return
	}
	_, err = database.AuthClient.VerifyIDToken(c, token)
	if err != nil {
		session := sessions.Default(c)
		refreshToken := session.Get("refresh_token").(string)
		if refreshToken != "" {
			resultInfo, err := controllers.RefreshToken(c, refreshToken)
			if err == nil {
				refreshToken = resultInfo["refresh_token"].(string)
				session.Set("refresh_token", refreshToken)
				session.Save()
				idToken := resultInfo["id_token"].(string)
				c.SetCookie("token", idToken, 3600, "/", "localhost", false, true)
				c.Next()
			}
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "data": nil})
		return
	}
	c.Next()
}

func NeedUserID(c *gin.Context) (userId string) {

	userId, err := c.Cookie("UID")
	if err != nil {
		token, _ := c.Cookie("token")
		userInfo, err := controllers.GetUserInfo(c, token)
		if err == nil {
			users, ok := userInfo["users"].([]interface{})
			if ok || len(users) > 0 {
				userData := users[0].(map[string]interface{})
				userId := userData["localId"].(string)
				return userId
			}
			return ""
		}
		return ""
	}
	return userId
}
