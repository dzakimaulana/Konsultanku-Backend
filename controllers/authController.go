package controllers

import (
	"fmt"
	"konsultanku-app/database"
	"konsultanku-app/errors"
	"konsultanku-app/models"
	"log"
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
	log.Printf("Successfully created user: %+v\n", createdUser)
	c.JSON(http.StatusOK, gin.H{"message": "Successfully create account", "data": createdUser.UID})
	return
}

func Login(c *gin.Context) {

	url := "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key="
	var user models.Login
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := map[string]interface{}{
		"email":             user.Email,
		"password":          user.Password,
		"returnSecureToken": true,
	}

	userData, err := SendRequest(c, url, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error", "data": nil})
		return
	}

	// set session
	session := sessions.Default(c)
	session.Set("refresh_token", userData["refreshToken"])
	session.Save()

	// set cookie
	token := userData["idToken"].(string)
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": userData})
	return

}

func Protected(c *gin.Context) {

	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "data": nil})
		return
	}

	data, err := database.AuthClient.VerifyIDToken(c, token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized", "data": nil})
		return
	}

	session := sessions.Default(c)
	session.Set("userId", data.UID)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
	return
}

func EmailVerification(c *gin.Context) {
	url := "https://identitytoolkit.googleapis.com/v1/accounts:sendOobCode?key="

	session := sessions.Default(c)
	idToken := session.Get("access_token")
	data := map[string]interface{}{
		"requestType": "PASSWORD_RESET",
		"idToken":     idToken.(string),
	}

	SendRequest(c, url, data)
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

	SendRequest(c, url, data)
	return
}

func Logout(c *gin.Context) {

	session := sessions.Default(c)
	userId := session.Get("userId")
	if userId == nil {
		fmt.Println("userId already gone")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Logout success", "data": nil})
		return
	}
	if err := database.AuthClient.RevokeRefreshTokens(c, userId.(string)); err != nil {
		errors.FirebaseAuthError(c, err)
		return
	}
	session.Clear()
	session.Save()
	c.JSON(200, gin.H{"message": "Logout sukses"})
}
