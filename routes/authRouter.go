package routes

import (
	controller "konsultanku-app/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRouter(incomingRoutes *gin.Engine) {

	authGroup := incomingRoutes.Group("/api/v1/auth")

	authGroup.POST("/register", controller.Register)
	authGroup.POST("/login", controller.Login)
	authGroup.POST("/email-verification", controller.EmailVerification)
	authGroup.POST("/reset-password", controller.ResetPassword)
	authGroup.POST("/logout", controller.Logout)

	// Jika Anda ingin menambahkan middleware ke seluruh grup, Anda bisa melakukannya di sini
	// authGroup.Use(middleware1, middleware2, ...)
}
