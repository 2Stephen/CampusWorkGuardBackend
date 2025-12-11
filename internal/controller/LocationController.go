package controller

import (
	"CampusWorkGuardBackend/internal/model/response"
	"CampusWorkGuardBackend/internal/service"
	"github.com/gin-gonic/gin"
)

func GetLocationController(ctx *gin.Context) {
	// 获取省市区列表
	keywords := ctx.Query("keywords")
	locations, err := service.GetLocationList(keywords)
	if err != nil {
		response.Fail(ctx, 500, "Failed to retrieve location list: "+err.Error())
		return
	}
	response.Success(ctx, locations)
}
