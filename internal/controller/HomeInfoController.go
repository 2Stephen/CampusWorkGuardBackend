package controller

import (
	"CampusWorkGuardBackend/internal/model/response"
	"CampusWorkGuardBackend/internal/service"
	"github.com/gin-gonic/gin"
)

func GetHomeStaticInfoController(ctx *gin.Context) {
	userId, exists := ctx.Get("userID")
	if !exists {
		response.Fail(ctx, 401, "用户未登录")
		return
	}
	role, exists := ctx.Get("role")
	if !exists {
		response.Fail(ctx, 401, "用户未登录")
		return
	}
	email, exists := ctx.Get("email")
	if !exists {
		response.Fail(ctx, 401, "用户未登录")
		return
	}

	info, err := service.GetHomeStaticInfo(userId.(int), role.(string), email.(string))
	if err != nil {
		response.Fail(ctx, 500, "获取个人基本信息失败: "+err.Error())
	}
	response.Success(ctx, info)
}
