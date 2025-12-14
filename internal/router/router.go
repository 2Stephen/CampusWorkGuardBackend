package router

import (
	controllers "CampusWorkGuardBackend/internal/controller"
	middlewares "CampusWorkGuardBackend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/uploads", "./uploads")
	// 允许跨域
	//r.Use(middlewares.CORSMiddleware())
	api := r.Group("/api")
	{
		api.GET("/school", controllers.GetSchoolListController)
		api.GET("/location", controllers.GetLocationController)
		// routes
		auth := api.Group("/auth")
		{
			student := auth.Group("/student")
			{
				student.POST("/register", controllers.AuthenticationStudentController)
				student.POST("/login", controllers.StudentLoginController)
				student.POST("/email_login", controllers.StudentEmailLoginController)
			}
			company := auth.Group("/company")
			{
				company.POST("/upload_license", controllers.UploadLicenseController)
				company.POST("/register", controllers.AuthenticationCompanyController)
				company.POST("/login", controllers.CompanyLoginController)
				company.POST("/email_login", controllers.CompanyEmailLoginController)
			}
			auth.POST("/send_code", controllers.SendCodeController)
		}
		studentUser := api.Group("/student_user")
		studentUser.Use(middlewares.TokenAuthRequired)
		{
			studentUser.POST("/submit", controllers.SubmitJobController)
			studentUser.POST("/set_password", controllers.SetStudentUserPasswordController)
			studentUser.GET("/profile_info", controllers.GetStudentUserProfileInfoController)
		}
		companyUser := api.Group("/company_user")
		companyUser.Use(middlewares.TokenAuthRequired)
		{
			companyUser.POST("set_password", controllers.SetCompanyUserPasswordController)
			companyUser.GET("delete", controllers.DeleteCompanyUserController)
			companyUser.GET("/profile_info", controllers.GetCompanyUserProfileInfoController)
			//companyUser.GET("/job_info", controllers.GetCompanyUserJobInfoController)
			companyUser.POST("/post_job", controllers.PostJobController)
		}
		home := api.Group("/home")
		home.Use(middlewares.TokenAuthRequired)
		{
			home.GET("/static_info", controllers.GetHomeStaticInfoController)
			home.POST("upload_avatar", controllers.UploadAvatarController)
		}
	}
	return r
}
