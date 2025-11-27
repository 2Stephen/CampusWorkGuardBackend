package router

import (
	controllers "CampusWorkGuardBackend/internal/controller"
	middlewares "CampusWorkGuardBackend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 允许跨域
	r.Use(middlewares.CORSMiddleware())
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
				student.POST("/login", controllers.StudentLoginController)
			}
			auth.POST("/send_code", controllers.SendCodeController)
		}
	}
	return r
}
