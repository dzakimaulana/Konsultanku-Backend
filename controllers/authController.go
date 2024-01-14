package controllers

import (
	"fmt"
	"konsultanku-app/database"
	"konsultanku-app/errors"
	"konsultanku-app/models"
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var person models.Register
	if err := c.ShouldBind(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid form data"})
		return
	}

	params := (&auth.UserToCreate{}).
		Email(person.Email).
		EmailVerified(false).
		PhoneNumber(person.PhoneNumber).
		Password(person.Password).
		DisplayName(person.Name).
		PhotoURL("http://www.example.com/12345678/photo.png").
		Disabled(false)
	createdUser, err := database.AuthClient.CreateUser(c, params)
	if err != nil {
		errors.FirebaseAuthError(c, err)
		return
	}

	result := map[string]interface{}{
		"uid":   createdUser.UID,
		"email": createdUser.Email,
		"name":  createdUser.DisplayName,
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully create account", "data": result})
	return

}

func Login(c *gin.Context) {

	url := "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key="
	var user models.Login
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Send request to Firebase
	jsonData := map[string]interface{}{
		"email":             user.Email,
		"password":          user.Password,
		"returnSecureToken": true,
	}
	dataUser, err := SendRequest(c, url, jsonData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error", "data": nil})
		return
	}

	// Get user info
	token := dataUser["idToken"].(string)
	userInfo, err := GetUserInfo(c, token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "data": nil})
		return
	}
	users := userInfo["users"].([]interface{})
	userData := users[0].(map[string]interface{})
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	if emailVerified := userData["emailVerified"].(bool); !emailVerified {
		c.JSON(http.StatusBadRequest, gin.H{"message": "verify your email first", "data": nil})
		return
	}

	// set session
	session := sessions.Default(c)
	session.Set("refresh_token", dataUser["refreshToken"])
	session.Save()

	// set cookie
	userId := userData["localId"].(string)
	c.SetCookie("UID", userId, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": userInfo})
	return

}

func EmailVerification(c *gin.Context) {

	url := "https://identitytoolkit.googleapis.com/v1/accounts:sendOobCode?key="
	idToken, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid token cookie"})
		return
	}
	data := map[string]interface{}{
		"requestType": "VERIFY_EMAIL",
		"idToken":     idToken,
	}

	userData, err := SendRequest(c, url, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Check your email", "data": userData["email"].(string)})
	return
}

func ResetPassword(c *gin.Context) {

	url := "https://identitytoolkit.googleapis.com/v1/accounts:sendOobCode?key="
	var user models.ResetPassword
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := map[string]interface{}{
		"requestType": "PASSWORD_RESET",
		"email":       user.Email,
	}

	userData, err := SendRequest(c, url, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Check your email", "data": userData["email"].(string)})
	return
}

func Logout(c *gin.Context) {

	uid, _ := c.Cookie("UID")
	if uid == "" {
		fmt.Println("userId already gone")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Logout success", "data": nil})
		return
	}
	if err := database.AuthClient.RevokeRefreshTokens(c, uid); err != nil {
		errors.FirebaseAuthError(c, err)
		// c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": nil})
		return
	}
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	cookies := c.Request.Cookies()
	for _, cookie := range cookies {
		c.SetCookie(cookie.Name, "", -1, "/", "", false, true)
	}

	c.Header("Cookie", "")
	c.JSON(http.StatusOK, gin.H{"message": "Logout success", "data": nil})
	return

}
