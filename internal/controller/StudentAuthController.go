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
	cHSIStudentInfo, err := service.StudentAuth(params)
	if err != nil {
		if err.Error() == "数据库保存学生信息失败" {
			response.Fail(c, 500, err.Error())
		} else {
			response.Fail(c, 403, err.Error())
		}
		return
	}
	// 返回认证结果
	if cHSIStudentInfo != nil {
		response.Success(c, nil)
	} else {
		response.Fail(c, 404, "学信网解析失败，请检查学信网验证码")
	}
}
