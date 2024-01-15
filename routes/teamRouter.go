package routes

import (
	controller "konsultanku-app/controllers"
	"konsultanku-app/middlewares"

	"github.com/gin-gonic/gin"
)

func TeamRouter(incomingRoutes *gin.Engine) {

	teamGroup := incomingRoutes.Group("/api/v1/team")
	teamGroup.Use(middlewares.StudentField, middlewares.TeamField)

	// incomingRoutes.POST("/api/v1/team/comment", controller.SendComment)
	teamGroup.GET("/api/v1/team/offers", controller.GetOffers)
	teamGroup.PUT("/api/v1/team/decision", controller.TeamDecision)
}
