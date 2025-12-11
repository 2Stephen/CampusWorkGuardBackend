package controller

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/model/response"
	"CampusWorkGuardBackend/internal/service"
	"github.com/gin-gonic/gin"
)

func SendCodeController(c *gin.Context) {
	var (
		params dto.SendCodeRequest
	)
	if err := c.ShouldBind(&params); err != nil {
		response.Fail(c, 400, "Invalid request parameters")
	}
	// 调用service发送验证码
	err := service.SendCode(params)
	if err != nil {
		if err.Error() == "用户不存在" || err.Error() == "用户已存在" {
			response.Fail(c, 403, err.Error())
		} else if err.Error() == "验证码发送频繁，请稍后重试" {
			response.Fail(c, 499, err.Error())
		} else {
			response.Fail(c, 500, "Failed to send code: "+err.Error())
		}
		return
	}
	response.Success(c, nil)
}
