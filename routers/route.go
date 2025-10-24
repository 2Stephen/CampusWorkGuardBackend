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
			auth.POST("/student", controllers.AuthenticationStudentController) // TODO: implement login handler
			auth.POST("/company", controllers.AuthenticationCompanyController) // TODO: implement register handler
		}
	}
	return r
}
