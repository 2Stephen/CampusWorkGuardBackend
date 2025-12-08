package controller

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/model/response"
	"CampusWorkGuardBackend/internal/service"
	"github.com/gin-gonic/gin"
	"log"
)

func GetSchoolListController(c *gin.Context) {
	search := c.Query("search")
	schools, err := service.GetSchoolList(search)
	if err != nil {
		log.Println("Error retrieving school list:", err)
		response.Fail(c, 500, "Failed to retrieve school list")
		return
	}
	response.Success(c, schools)
}

func AuthenticationStudentController(c *gin.Context) {
	var (
		params dto.StudentAuthParams
	)
	if err := c.ShouldBind(&params); err != nil {
		log.Println("Error binding request parameters:", err)
		c.JSON(400, gin.H{"error": "Invalid request parameters"})
		// c.Abort()
		return
	}
	// 调用service进行认证逻辑处理
	cHSIStudentInfo, token, err := service.StudentAuth(params)
	if err != nil {
		if err.Error() == "数据库保存学生信息失败" || err.Error() == "获取邮箱验证码失败" {
			response.Fail(c, 500, err.Error())
		} else {
			response.Fail(c, 403, err.Error())
		}
		return
	}
	// 返回认证结果
	if cHSIStudentInfo != nil {
		response.Success(c, gin.H{"token": token})
	} else {
		response.Fail(c, 404, "学信网解析失败，请检查学信网验证码")
	}
}

func StudentLoginController(c *gin.Context) {
	var (
		params dto.StudentLoginParams
	)
	if err := c.ShouldBind(&params); err != nil {
		log.Println("Error binding request parameters:", err)
		response.Fail(c, 400, "Invalid request parameters")
		return
	}
	// 调用service进行登录逻辑处理
	token, err := service.StudentLogin(params)
	if err != nil {
		if err.Error() == "学生登录失败，检查学号或密码是否正确" || err.Error() == "用户未设置密码，请使用邮箱验证登录后设置密码" {
			response.Fail(c, 403, err.Error())
			return
		}
		response.Fail(c, 500, "Failed to login: "+err.Error())
		return
	}
	response.Success(c, gin.H{"token": token})
}

func StudentEmailLoginController(c *gin.Context) {
	var (
		params dto.StudentEmailLoginParams
	)
	if err := c.ShouldBind(&params); err != nil {
		log.Println("Error binding request parameters:", err)
		response.Fail(c, 400, "Invalid request parameters")
		return
	}
	// 调用service进行登录逻辑处理
	token, err := service.StudentEmailLogin(params)
	if err != nil {
		if err.Error() == "邮箱验证码有误" || err.Error() == "邮箱验证码已过期，请重新获取" {
			response.Fail(c, 403, err.Error())
			return
		}
		response.Fail(c, 500, "Failed to login: "+err.Error())
		return
	}
	response.Success(c, gin.H{"token": token})
}
