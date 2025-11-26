package router

import (
	controllers "CampusWorkGuardBackend/internal/controller"
	middlewares "CampusWorkGuardBackend/internal/middleware/TokenAuthRequired"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	api.Use(middlewares.TokenAuthRequired())
	{
		api.GET("/school", controllers.GetSchoolListController)
		// routes
		auth := api.Group("/auth")
		{
			student := auth.Group("/student")
			{
				student.POST("/register", controllers.AuthenticationStudentController)
			}
		}
	}
	return r
}
