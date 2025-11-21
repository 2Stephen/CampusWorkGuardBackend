package routers

import (
	"CampusWorkGuardBackend/controllers/AuthenticationModuleController"
	"CampusWorkGuardBackend/middlewares/TokenAuthRequired"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	api.Use(middlewares.TokenAuthRequired())
	{
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
