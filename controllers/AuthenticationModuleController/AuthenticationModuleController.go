package controllers

import (
	middlewares "CampusWorkGuardBackend/middlewares/AuthenticationModule"
	"CampusWorkGuardBackend/models/request"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func AuthenticationStudentController(c *gin.Context) {
	var (
		params request.StudentAuthParams
	)
	if err := c.ShouldBind(&params); err != nil {
		log.Println("Error binding request parameters:", err)
		c.JSON(400, gin.H{"error": "Invalid request parameters"})
		// c.Abort()
		return
	}
	// 调用中间件进行认证逻辑处理
	cHSIStudentInfo := &middlewares.CHSIStudentInfo{}
	result := cHSIStudentInfo.StudentAuth(params)
	// 返回认证结果
	if result {
		// 输出认证成功信息
		fmt.Printf("Student authentication successful: %+v\n", cHSIStudentInfo)
		c.JSON(200, gin.H{"message": "Student authentication successful"})
	} else {
		c.JSON(401, gin.H{"message": "Student authentication failed"})
	}
}
func AuthenticationCompanyController(c *gin.Context) {

}
