package routes

import (
	controller "konsultanku-app/controllers"
	"konsultanku-app/middlewares"

	"github.com/gin-gonic/gin"
)

func StudentRouter(incomingRoutes *gin.Engine) {

	studentGroup := incomingRoutes.Group("/api/v1/student")
	studentGroup.Use(middlewares.StudentOnly)

	studentGroup.POST("/registration", controller.CreateStudentProfile)
	studentGroup.GET("/problems", controller.Problem)
	studentGroup.GET("/problems/:id", controller.ProblemByID)

}
