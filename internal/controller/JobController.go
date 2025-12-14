package controller

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/model/response"
	"CampusWorkGuardBackend/internal/service"
	"github.com/gin-gonic/gin"
)

func GetCompanyUserJobInfoController(c *gin.Context) {
	// TODO: Implement the logic to get company user job info
}

func PostJobController(c *gin.Context) {
	var (
		params dto.PostJobParams
	)
	if err := c.ShouldBind(&params); err != nil {
		response.Fail(c, 400, "Invalid request parameters")
	}
	userID, exists := c.Get("userID")
	if !exists {
		response.Fail(c, 401, "用户未认证")
		return
	}
	email, exists := c.Get("email")
	if !exists {
		response.Fail(c, 401, "用户未认证")
		return
	}
	err := service.PostJobService(params, userID.(int), email.(string))
	if err != nil {
		if err.Error() == "企业用户不存在" || err.Error() == "企业用户未通过认证，无法发布职位" || err.Error() == "用户邮箱与认证邮箱不匹配" {
			response.Fail(c, 403, err.Error())
			return
		}
		response.Fail(c, 500, "Failed to post job: "+err.Error())
		return
	}
	response.Success(c, nil)
}
