package controller

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/model/response"
	"CampusWorkGuardBackend/internal/service"
	"github.com/gin-gonic/gin"
)

func GetAdminCompanyListController(c *gin.Context) {
	var (
		params dto.CompanyListRequest
	)
	if err := c.ShouldBind(&params); err != nil {
		response.Fail(c, 400, "参数绑定失败"+err.Error())
		return
	}
	CompanyList, total, err := service.GetAdminCompanyListService(&params)
	if err != nil {
		response.Fail(c, 500, "获取公司列表失败"+err.Error())
		return
	}
	response.Success(c, gin.H{
		"companys": CompanyList,
		"total":    total,
	})
}
