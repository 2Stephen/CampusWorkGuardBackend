package controller

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/model/response"
	"CampusWorkGuardBackend/internal/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func SetCompanyUserPasswordController(c *gin.Context) {
	// 调用service进行设置密码逻辑处理
	var (
		params dto.SetCompanyUserPasswordParams
	)
	if err := c.ShouldBind(&params); err != nil {
		response.Fail(c, 400, "Invalid request parameters")
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		response.Fail(c, 401, "用户未认证")
		return
	}
	err := service.SetCompanyUserPassword(params, strconv.Itoa(userID.(int)))
	if err != nil {
		if err.Error() == "密码长度不足，至少需要8位" || err.Error() == "密码必须包含字母和数字" || err.Error() == "密码长度过长，不能超过64位" {
			response.Fail(c, 403, err.Error())
		} else {
			response.Fail(c, 500, "Failed to set password: "+err.Error())
		}
		return
	}
	response.Success(c, nil)
}
