package routes

import (
	controller "konsultanku-app/controllers"
	"konsultanku-app/middlewares"

	"github.com/gin-gonic/gin"
)

func TeamRouter(incomingRoutes *gin.Engine) {

	teamGroup := incomingRoutes.Group("/api/v1/team")
	teamGroup.Use(middlewares.StudentField, middlewares.TeamField)

	teamGroup.POST("/comment/:problemID", middlewares.IsLeader, controller.AddComment)
	teamGroup.GET("/offers", controller.GetOffers)
	teamGroup.PUT("/decision", controller.TeamDecision)
}
