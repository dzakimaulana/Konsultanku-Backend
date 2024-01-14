package routes

import (
	controller "konsultanku-app/controllers"
	"konsultanku-app/middlewares"

	"github.com/gin-gonic/gin"
)

func MseRouter(incomingRoutes *gin.Engine) {

	mseGroup := incomingRoutes.Group("/api/v1/mse")
	mseGroup.Use(middlewares.MseField)

	mseGroup.POST("/registration", controller.CreateMseProfile)
	mseGroup.POST("/create-problem", controller.CreateProblem)
	mseGroup.GET("/comments", controller.AllComments)
	mseGroup.POST("/send-offer/:id", controller.SendOffer)
}
