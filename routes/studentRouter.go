package routes

import (
	controller "konsultanku-app/controllers"
	"konsultanku-app/middlewares"

	"github.com/gin-gonic/gin"
)

func StudentRouter(incomingRoutes *gin.Engine) {

	studentGroup := incomingRoutes.Group("/api/v1/student")
	studentGroup.Use(middlewares.Authenticate, middlewares.StudentField)

	studentGroup.POST("/registration", controller.CreateStudentProfile)
	studentGroup.POST("/create-team", controller.BuildTeam)
	studentGroup.PUT("/profile", controller.UpdateStudentProfile)
	studentGroup.GET("/problems", controller.Problem)
	studentGroup.GET("/problems/:id", controller.ProblemByID)

}
