package controller

import (
	"CampusWorkGuardBackend/internal/dto"
	"CampusWorkGuardBackend/internal/model/response"
	"CampusWorkGuardBackend/internal/service"
	"github.com/gin-gonic/gin"
	"log"
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

func ReviewCompanyController(c *gin.Context) {
	var (
		params dto.CompanyReviewRequest
	)
	if err := c.ShouldBind(&params); err != nil {
		response.Fail(c, 400, "参数绑定失败"+err.Error())
		return
	}
	if err := service.ReviewCompanyService(&params); err != nil {
		log.Println(err)
		response.Fail(c, 500, "审核公司失败"+err.Error())
		return
	}
	response.Success(c, nil)
}
