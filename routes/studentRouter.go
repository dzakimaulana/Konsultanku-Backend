package routes

import (
	controller "konsultanku-app/controllers"
	"konsultanku-app/middlewares"

	"github.com/gin-gonic/gin"
)

func StudentRouter(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("/api/v1/student/registration", controller.CreateStudentProfile)

	studentGroup := incomingRoutes.Group("/api/v1/student")
	studentGroup.Use(middlewares.Authenticate, middlewares.StudentField)

	studentGroup.POST("/create-team", controller.BuildTeam)
	studentGroup.PUT("/profile", controller.UpdateStudentProfile)
	studentGroup.PUT("/join-team", controller.JoinTeam)
	studentGroup.PUT("/like", controller.LikeProblem)
	studentGroup.GET("/problems", controller.Problem)
	studentGroup.GET("/problems/:id", controller.ProblemByID)

}
